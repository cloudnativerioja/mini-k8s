# Mini cluster de kubernetes

## Requisitos

- Docker
- Go
- Cliente de kubernetes (kubectl, lens, k9s, etc)

## Despliegue

* Para desplegar el cluster, ejecutar el siguiente comando:

```bash
make run ACTION=create
```

* Para destruir el cluster, ejecutar el siguiente comando:

```bash
make run ACTION=delete
```
## Configuración del cluster

Dentro de la carpeta config hay un yaml con la configuración de nodos y demás opciones. Ejemplo de configuración https://k3d.io/v5.0.1/usage/configfile/
