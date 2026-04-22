#!/bin/bash

function ctrl_c(){
    echo -e "\n[!] Saliendo...\n"
    tput cnorm 2>/dev/null
    exit 1
}

trap ctrl_c SIGINT

set -e

BIN_NAME="emlcheck"
INSTALL_PATH="/usr/local/bin/$BIN_NAME"

echo "[+] Compiling $BIN_NAME..."
go build -o "$BIN_NAME" main.go

if [ ! -f "$BIN_NAME" ]; then
    echo "[!] Error: No se generó el binario"
    exit 1
fi

echo "[+] Instalando $BIN_NAME en /usr/local/bin"
sudo mkdir -p /usr/local/bin
sudo install -m 755 "$BIN_NAME" "$INSTALL_PATH"

if [ -x "$INSTALL_PATH" ]; then
    echo "[+] Instalación completada"
    echo "[i] Binario: $INSTALL_PATH"
else
    echo "[!] Error: No se pudo instalar correctamente"
    exit 1
fi

if command -v "$BIN_NAME" >/dev/null 2>&1; then
    echo "[+] Disponible en PATH"
else
    echo "[!] El binario se instaló, pero /usr/local/bin no está en tu PATH actual"
    echo "[i] Prueba ejecutarlo así: $INSTALL_PATH"
fi