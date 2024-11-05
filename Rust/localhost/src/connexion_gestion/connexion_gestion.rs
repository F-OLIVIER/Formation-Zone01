#[warn(unused_imports)]
use crate::handle::handle_client::*;
use crate::MAX_EVENTS;
use core::str;
use std::process::Command;
use std::net::TcpListener;
use std::os::unix::io::AsRawFd;

#[cfg(target_os = "macos")]
pub fn run_kqueue_macos(listeners: Vec<TcpListener>) {
    use libc::{c_int, kevent, kqueue, EVFILT_READ, EV_ADD, EV_ENABLE, EV_CLEAR};
    
    // Création d'un nouvel kqueue
    let kq = unsafe { kqueue() };
    if kq == -1 {
        eprintln!("Erreur lors de la création de kqueue");
        return;
    }

    let mut changes = vec![];
    let mut listener_fds = vec![];

    // Enregistrement de tous les listeners dans kqueue
    for listener in &listeners {
        let listener_fd = listener.as_raw_fd();
        listener_fds.push(listener_fd);
        let change = libc::kevent {
            ident: listener_fd as usize,
            filter: EVFILT_READ,
            flags: EV_ADD | EV_ENABLE | EV_CLEAR,
            fflags: 0,
            data: 0,
            udata: std::ptr::null_mut(),
        };
        changes.push(change);
    }

    // Enregistrer les changements dans kqueue
    unsafe {
        kevent(
            kq,
            changes.as_mut_ptr(),
            changes.len() as c_int,
            std::ptr::null_mut(),
            0,
            std::ptr::null_mut(),
        );
    }

    // Boucle principale pour traiter les événements
    loop {
        let mut events = vec![libc::kevent {
            ident: 0,
            filter: 0,
            flags: 0,
            fflags: 0,
            data: 0,
            udata: std::ptr::null_mut(),
        }; MAX_EVENTS];

        let nev = unsafe {
            kevent(
                kq,
                std::ptr::null_mut(),
                0,
                events.as_mut_ptr(),
                MAX_EVENTS as c_int,
                std::ptr::null_mut(),
            )
        };

        // Vérification de l'appel à kevent
        if nev < 0 {
            eprintln!("Erreur lors de l'appel à kevent");
            break;
        }

        // Traitement des événements retournés
        for i in 0..nev as usize {
            let event = &events[i];

            // Si l'événement provient de l'un des listeners, accepter la connexion
            if listener_fds.contains(&(event.ident as i32)) {
                let index = listener_fds.iter().position(|&fd| fd == event.ident as i32).unwrap();
                let listener = &listeners[index];

                // Tente d'accepter une connexion
                match listener.accept() {
                    Ok((mut stream, addr)) => {
                        println!("Connexion acceptée de : {:?}", addr);
                        handle_client(&mut stream);
                    }
                    Err(e) => eprintln!("Erreur lors de l'acceptation de la connexion : {}", e),
                }
            }
        }
    }

    // Fermeture du descripteur de kqueue
    unsafe {
        libc::close(kq);
    }
}

#[cfg(target_os = "linux")]
pub fn run_epoll_linux(listeners: Vec<TcpListener>) {
    use libc::{c_int, epoll_create1, epoll_ctl, epoll_event, epoll_wait, EPOLLIN, EPOLL_CTL_ADD};

    // Création d'un nouvel epoll
    let epoll_fd = unsafe { epoll_create1(0) };
    if epoll_fd == -1 {
        eprintln!("Erreur lors de la création d'epoll");
        return;
    }

    let mut events = vec![epoll_event { events: 0, u64: 0 }; MAX_EVENTS];

    // Enregistrement de tous les listeners dans epoll
    for listener in &listeners {
        let listener_fd = listener.as_raw_fd();
        let mut event = epoll_event {
            events: EPOLLIN as u32,
            u64: listener_fd as u64,
        };
        if unsafe { epoll_ctl(epoll_fd, EPOLL_CTL_ADD, listener_fd, &mut event) } == -1 {
            eprintln!("Erreur lors de l'enregistrement de l'événement dans epoll");
            return;
        }
    }

    loop {
        let nev = unsafe { epoll_wait(epoll_fd, events.as_mut_ptr(), MAX_EVENTS as c_int, -1) };
        if nev < 0 {
            eprintln!("Erreur lors de l'appel à epoll_wait");
            break;
        }

        // Traitement des événements retournés
        for i in 0..nev as usize {
            let event = &events[i];

            // Si l'événement provient de l'un des listeners, accepter la connexion
            for listener in &listeners {
                let listener_fd = listener.as_raw_fd();
                if event.u64 == listener_fd as u64 {
                    match listener.accept() {
                        Ok((mut stream, addr)) => {
                            println!("Connexion acceptée de : {:?}", addr);
                            handle_client(&mut stream);
                        }
                        Err(e) => eprintln!("Erreur lors de l'acceptation de la connexion : {}", e),
                    }
                }
            }
        }
    }

    unsafe {
        libc::close(epoll_fd);
    }
}

pub fn path_server() -> String {
    if cfg!(target_os = "macos") || cfg!(target_os = "linux") {
        // Exécute la commande 'pwd' pour récupérer le chemin du répertoire actuel
        match Command::new("pwd").output() {
            Ok(output) => {
                if output.status.success() {
                    // Convertir le résultat en chaîne de caractères UTF-8, avec gestion d'erreur
                    match str::from_utf8(&output.stdout) {
                        Ok(path) => path.trim().to_string(),
                        Err(e) => {
                            eprintln!("Error converting output to UTF-8: {}", e);
                            String::new() // Retourne une chaîne vide
                        }
                    }
                } else {
                    eprintln!(
                        "Failed to get current directory path. Command returned non-zero status."
                    );
                    String::new() // Retourne une chaîne vide
                }
            }
            Err(e) => {
                eprintln!("Failed to execute 'pwd' command: {}", e);
                String::new() // Retourne une chaîne vide
            }
        }
    } else {
        eprintln!("Unsupported operating system");
        String::new() // Par défaut, retourne une chaîne vide si le système n'est pas pris en charge
    }
}