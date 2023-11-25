# M-Authentication 🚀

¡Bienvenido a M-Authentication, tu solución todo en uno para la autenticación de usuarios! ✨ Este servicio está construido en Go con gin, diseñado para simplificar el proceso de autenticación y eliminar la necesidad de crear microservicios adicionales cada año.

## Ventajas 🌟

- **Simplicidad:** Olvídate de configurar múltiples microservicios para autenticar usuarios. ¡M-Authentication lo hace todo por ti!
- **Eficiencia:** Despliega rápidamente tu sistema de autenticación sin perder tiempo en configuraciones complicadas.
- **Adaptabilidad:** Compatible con diversas estrategias de autenticación para satisfacer tus necesidades específicas.

## Configuración 🛠️

### Configuración del Servidor (`config/settings.xml`)
Añade tu configuración de servidor en `config/settings.xml`  siguiendo las reglas XML proporcionadas.

```xml
<!-- Ejemplo de config/settings.xml -->
<Config>
  <service>
    <authMethod output="consola" type="basic" />
  </service>
  <server>
    <port>8080</port>
  </server>
</Config>
```
### Configuración del Servidor (`config/settings.xml`)

Define las reglas de autenticación básica en `config/auth/basic.xml` según el esquema XML indicado.
```xml
<!-- Ejemplo de config/auth/basic.xml -->
<Config>
  <connection id="1" type="mysql" host="localhost" port="3306" database="usuarios" user="admin" password="admin123" />
  <auth useRoles="true" routerName="authRouter">
    <table name="usuarios">
      <user column="username" />
      <password column="password">
        <encrypt algorithm="HS384" source="local" key="mySecretKey" />
      </password>
    </table>
    <roles>
      <global>
        <claims>
          <DataSource type="mysql" name="roles" column="rol" />
        </claims>
      </global>
      <role name="admin">
        <claims>
          <DataSource type="mysql" name="admin_claims" column="claim_name" />
        </claims>
      </role>
    </roles>
  </auth>
</Config>
```

**¡Listo para usar!** 🚀

## Instalación 📦

1. Clona este repositorio:

   ```bash
   git clone https://github.com/mrthoabby/m-authentication
   ```
2. Instala las dependencias:

   ```bash
    go get -u github.com/gin-gonic/gin

   ```
3. Ejecuta el servicio:

   ```bash
   go run main.go

   ```

#Contribuciones 🤝
¡Este proyecto es de código abierto y está bajo licencia MIT! 👩‍💻 Siéntete libre de hacer un fork y contribuir para hacerlo aún mejor.

¡Sé creativo y diviértete con M-Authentication! 🚀🔐