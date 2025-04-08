use local_ip_address::list_afinet_netifas;
use std::io::{self, Error};
use std::net::{IpAddr, Ipv4Addr};
use std::process::Command;
use std::str::FromStr;

pub fn get_client_ip_local() -> Result<IpAddr, Error> {
    // Appel de la fonction pour lister les interfaces réseau
    let result = list_afinet_netifas();

    // Gestion explicite de l'erreur
    let interfaces = match result {
        Ok(ifaces) => ifaces, // Si l'appel réussit, obtenir la liste des interfaces
        Err(e) => return Err(io::Error::new(io::ErrorKind::Other, e)), // Propager l'erreur
    };

    // Sélectionner l'interface correcte en fonction du système d'exploitation
    #[cfg(target_os = "linux")]
    {
        for (iface, ip) in interfaces {
            println!("iface : {} : {}", iface, ip);
            if iface.starts_with("eth") || iface.starts_with("en") || iface.starts_with("wlp") {
                return Ok(ip);
            }
        }
    }

    #[cfg(target_os = "macos")]
    {
        for (iface, ip) in interfaces {
            if iface.starts_with("en") || iface.starts_with("en0") {
                return Ok(ip);
            }
        }
    }

    #[cfg(target_os = "windows")]
    {
        for (iface, ip) in interfaces {
            if iface.starts_with("Ethernet") || iface.starts_with("Wi-Fi") {
                return Ok(ip);
            }
        }
    }

    // Si aucune adresse IP n'a été trouvée, renvoyer une adresse localhost par défaut
    Ok(IpAddr::V4(Ipv4Addr::LOCALHOST))
}

// Fonction pour obtenir l'adresse IP publique
pub fn get_client_ip_public() -> Result<IpAddr, Box<dyn std::error::Error>> {
    let curl_check = Command::new("curl").arg("--version").output();

    match curl_check {
        Ok(_) => {
            let output = Command::new("curl")
                .arg("-s")
                .arg("https://api64.ipify.org")
                .output()?;

            let ip_str = String::from_utf8(output.stdout)?.trim().to_string();
            let ip: IpAddr = IpAddr::from_str(&ip_str)?; // Utilisation de IpAddr::from_str
            Ok(ip)
        }
        Err(_) => {
            Err(Box::new(std::io::Error::new(std::io::ErrorKind::NotFound, "curl n'est pas disponible sur ce système")))
        }
    }
}
