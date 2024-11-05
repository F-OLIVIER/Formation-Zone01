#!/bin/bash

# Lecture du fichier ./src/config/config.txt
while IFS=' ' read -r ip_port domain_name; do
    # Vérifie si le nom de domaine est non vide
    echo "Ligne lue : '$ip_port $domain_name'"

    if [[ -n "$domain_name" ]]; then
        # Extraction de l'adresse IP et du port
        ip="${ip_port%:*}"        # Tout avant le dernier ":"
        port="${ip_port##*:}"     # Tout après le dernier ":"
        
        echo "IP : $ip, Port : $port, Domain : $domain_name"

        # Vérifie si la ligne existe déjà dans /etc/hosts
        if ! grep -q "$ip[[:space:]]\+$domain_name" /etc/hosts; then
            # Ajout aux tableaux uniquement si ce n'est pas déjà présent
            IP_ADDRESSES+=("$ip")
            PORTS+=("$port")
            DOMAIN_NAMES+=("$domain_name")

            # Ajout dans le fichier /etc/hosts
            # echo "$ip $domain_name" | sudo tee -a /etc/hosts
            echo "Ajouté: $ip $domain_name"
        else
            echo "Déjà présent: $ip $domain_name"
        fi
    fi
done < ./src/config/config.txt

# Affichage des tableaux pour vérification
# echo "DOMAIN_NAMES=(${DOMAIN_NAMES[*]})"
# echo "IP_ADDRESSES=(${IP_ADDRESSES[*]})"
# echo "PORTS=(${PORTS[*]})"

# Détection du système d'exploitation
if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Système détecté : MacOS"

    echo "Ajout des entrées dans le fichier hosts..."
    for i in "${!DOMAIN_NAMES[@]}"; do
        IP=${IP_ADDRESSES[i]}
        DOMAIN=${DOMAIN_NAMES[i]}
        echo "$IP $DOMAIN" | sudo tee -a /etc/hosts
    done

    for IP in "${IP_ADDRESSES[@]}"; do
        echo "Ajout de l'alias $IP sur l'interface loopback..."
        sudo ifconfig lo0 alias $IP
    done

elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    echo "Système détecté : Ubuntu"

    echo "Ajout des entrées dans le fichier hosts..."
    for i in "${!DOMAIN_NAMES[@]}"; do
        IP=${IP_ADDRESSES[i]}
        DOMAIN=${DOMAIN_NAMES[i]}
        echo "$IP $DOMAIN" | sudo tee -a /etc/hosts
    done

    for IP in "${IP_ADDRESSES[@]}"; do
        echo "Ajout de l'alias $IP sur l'interface loopback..."
        sudo ip addr add $IP dev lo
    done

else
    echo "Système d'exploitation non pris en charge."
    exit 1
fi
