NOMBRE_ARCHIVO := "gomountain"
CARPETA_BINARIOS := "./bin"
CARPETA_CODIGO := "./internal"

# Construcción del proyecto
build:
    echo "Construyendo el proyecto"
    mkdir -p {{CARPETA_BINARIOS}}
    go build -o {{CARPETA_BINARIOS}}/{{NOMBRE_ARCHIVO}} {{CARPETA_CODIGO}}

# Instalar dependencias
install_deps:
    echo "Instalando dependencias"
    go mod tidy

# Eliminar binarios
clean:
    echo "Elminando los binarios"
    rm -f {{CARPETA_BINARIOS}}/{{NOMBRE_ARCHIVO}}
    go clean {{CARPETA_CODIGO}}

# Comprobación sintaxis del proyecto
check:
    echo "Comprobando sintaxis del proyecto"
    gofmt -e {{CARPETA_CODIGO}} > /dev/null
