## **Laboratorio Nº 07**

**E. P.:** Ingeniería de Sistemas (plan 2018 R) **Asignatura:** IS489 Pruebas y Aseguramiento de Calidad de Software **Docente:** Mtro. Ing. Oscar Meneses Yaranga **Semestre Académico** : 2026-I Presencial **SESIÓN 07: EXAMEN PARCIAL**

## **I. CASOS PROPUESTOS:**

Revisar lo siguiente:

## **CASO I: Elaboración de un plan de pruebas para un sistema académico o empresarial** .

**Título del caso** : Elaboración de un plan de pruebas de software para …

**Objetivo** : Definir el alcance, los objetivos, los recursos, el cronograma y la estrategia para garantizar la **calidad de un producto Software** .

## **Instrucciones:**

**1. Identificar el Alcance y los Objetivos**

- **Alcance:** Define claramente qué módulos o funcionalidades se probarán y cuáles quedarán fuera (para evitar confusiones).

- **Objetivos:** Establece qué esperas lograr con las pruebas (por ejemplo: validar que el sistema soporte 5,000 usuarios concurrentes o asegurar el 100 % de cobertura en el flujo de inicio de sesión).

**2. Definir la Estrategia de Pruebas**

Establece los métodos y marcos de trabajo (frameworks como JUnit/pytest) que utilizará el equipo. Debes incluir:

- **Tipos de pruebas:** Pruebas unitarias, de integración, de sistema, de regresión, de rendimiento o de seguridad.

- **Enfoque:** ¿Serán pruebas manuales, automatizadas o ambas?

**3. Asignar Recursos y Entornos**

- **Equipo:** Define los roles necesarios (QA, desarrolladores, analistas, etc.) y las responsabilidades de cada uno.

- **Entornos:** Describe la infraestructura, servidores, dispositivos y herramientas (como Jira o TestRail) que el equipo utilizará.

**4. Establecer el Cronograma**

Distribuye las actividades de testing dentro de los hitos del proyecto o _sprints_ (si trabajas con metodologías ágiles). Asegúrate de asignar tiempos específicos para la ejecución y la corrección de errores.

**5. Criterios de Entrada y Salida**

- **Criterio de entrada:** Qué condiciones deben cumplirse para iniciar las pruebas (por ejemplo: que el código pase la prueba de humo o _smoke test_ y el entorno esté configurado).

- **Criterio de salida:** Cuándo se consideran finalizadas las pruebas (por ejemplo: 95 % de los casos de prueba superados y sin defectos críticos abiertos).

**6. Gestión de Riesgos y Mitigación**

Identifica posibles obstáculos, como retrasos en la entrega de código, falta de datos de prueba o cambios en los requisitos. Diseña un plan de contingencia para cada uno que se identifique.

## **7. Entregables**

Detalla la documentación que se generará, como los casos de prueba, los informes de errores ( _bugs_ ) y el informe final de cierre de pruebas.

**Ejemplo** : Para un sistema integral (ERP/CRM) que puede ser completamente nuevo, el enfoque sería centrarse en asegurar que los módulos individuales funcionen correctamente y que la **integración de datos** entre ellos sea fluida.

## **Plan de pruebas de tu software empresarial …**

## **1. Alcance del Proyecto**

- **Módulos a probar:** Ventas (cotizaciones, facturas), Marketing (campañas, leads), RRHH (nómina, asistencia) y Postventa (tickets de soporte).

- **Integraciones clave:** Traspaso de un lead (Marketing) a cliente (Ventas), y de cliente a soporte (Postventa).

- **Fuera de alcance:** Migración de datos históricos (al ser un proyecto nuevo, se inicia con base de datos limpia).

## **2. Estrategia y Tipos de Prueba Prioritarios**

- **Pruebas Funcionales:** Validar que cada regla de negocio se cumpla (ej: cálculo de impuestos en Ventas o descuentos en Marketing).

- **Pruebas de Integración:** Verificar que la información viaje correctamente entre los cuatro módulos sin duplicarse ni perderse.

- **Pruebas de Roles y Permisos (Seguridad):** Crucial para empresas. Asegurar que un usuario de Marketing no pueda ver la nómina de RRHH.

- **Automatización:** Priorizar la automatización en el módulo de Ventas y Nómina (RRHH) por ser flujos críticos y repetitivos.

## **3. Entorno y Datos de Prueba**

- **Entorno de QA:** Un servidor aislado idéntico al de producción, pero con datos ficticios.

- **Datos requeridos:** Listas de empleados simulados, catálogos de productos cargados y clientes de prueba para ejecutar los flujos completos.

## **4. Criterios de Éxito (Entrada y Salida)**

- **Entrada:** Código desplegado en QA, base de datos inicial configurada y HU (Historias de Usuario) documentadas.

- **Salida:** 100% de las pruebas críticas de Ventas y RRHH aprobadas, y cero defectos bloqueantes o críticos activos.

## **5. Riesgos Específicos del Proyecto**

- **Riesgo:** Complejidad en las dependencias de datos entre módulos.

- **Mitigación:** Diseñar los casos de prueba siguiendo el ciclo de vida del cliente (Lead → Venta → Postventa).

## **Entregables:**

## **E1. Tabla Detallada de Casos de Prueba (E2E e Intermodulares)**

| **ID**          | **Módulo**<br>**/**<br>**Componente** | **/**<br>**onente**<br>**Descripción**<br>**del**<br>**Caso de Prueba**         | **del**<br>**Precondiciones**                                                                      | **Pasos**<br>**para**<br>**la**<br>**Ejecución**                                                                                      | **la**<br>**Resultado Esperado Prioridad**                                                                                                                                    | **Resultado Esperado Prioridad**                                                          |
| --------------- | ------------------------------------- | ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| **CP-01**       | Marketing<br>a<br>Ventas              | a<br>Convertir un Lead<br>calificado<br>en<br>Oportunidad<br>de<br>Venta.       | Convertir un Lead<br>en<br>de<br>Lead<br>creado<br>con<br>estado "Calificado".                     | con<br>1.<br>Ir<br>a<br>Marketing.<br>2.<br>Seleccionar<br>Lead.<br>3. Hacer clic en "Convertir<br>a Cliente".                        | Marketing.<br>Lead.<br>3. Hacer clic en "Convertir<br>El Lead desaparece de<br>Marketing y se crea el<br>Cliente y la Oportunidad<br>en Ventas.                               | El Lead desaparece de<br>Marketing y se crea el<br>Cliente y la Oportunidad<br>**Alta**   |
| **CP-02**Ventas | Ventas                                | Generar<br>factura<br>y<br>calcular<br>impuestos<br>correctamente.              | y<br>impuestos<br>Cliente y producto<br>con<br>stock<br>registrados.                               | Cliente y producto<br>stock<br>1.<br>Crear<br>cotización.<br>2.<br>Aprobar<br>cotización.<br>3.<br>Generar<br>factura<br>electrónica. | cotización.<br>cotización.<br>factura<br>Factura generada con<br>ID único. El cálculo de<br>impuestos coincide con<br>la leylocal.                                            | Factura generada con<br>ID único. El cálculo de<br>impuestos coincide con<br>**Alta**     |
| **CP-03**RRHH   | RRHH                                  | Procesar<br>nómina<br>mensual con horas<br>extra.                               | nómina<br>mensual con horas<br>Empleado<br>activo<br>con<br>registro<br>de<br>asistencia.          | activo<br>de<br>1.<br>Ir<br>a<br>RRHH.<br>2. Seleccionar empleado.<br>3. Importar horas extra.<br>4. Calcular nómina.                 | RRHH.<br>2. Seleccionar empleado.<br>3. Importar horas extra.<br>El sistema genera el<br>recibo<br>de<br>pago<br>sumando el sueldo base<br>más las horas extra<br>calculadas. | El sistema genera el<br>pago<br>sumando el sueldo base<br>más las horas extra<br>**Alta** |
| **CP-04**       | Ventas<br>a<br>Postventa              | a<br>Crear<br>ticket<br>de<br>soporte asociado a<br>un<br>producto<br>comprado. | de<br>soporte asociado a<br>producto<br>Factura<br>de venta<br>cerrada y pagada.                   | de venta<br>1.<br>Ir<br>a<br>Postventa.<br>2. Crear nuevo ticket.<br>3. Buscar cliente por ID.<br>4. Asociar producto.                | Postventa.<br>2. Crear nuevo ticket.<br>3. Buscar cliente por ID.<br>El<br>ticket<br>se<br>guarda<br>vinculando<br>correctamente<br>el<br>historial de compra del<br>cliente. | guarda<br>el<br>historial de compra del<br>Media                                          |
| **CP-05**       | Seguridad<br>/<br>Roles               | /<br>Restringir acceso de<br>usuario de Marketing<br>a RRHH.                    | Restringir acceso de<br>usuario de Marketing<br>Usuario<br>con<br>rol<br>"Ejecutivo<br>Marketing". | rol<br>1.<br>Iniciar<br>sesión<br>con<br>usuario<br>de<br>Marketing.<br>2. Intentar ingresar a URL<br>de Nómina.                      | con<br>Marketing.<br>2. Intentar ingresar a URL<br>El sistema bloquea el<br>acceso y muestra un<br>error<br>de<br>"Permisos<br>insuficientes".                                | El sistema bloquea el<br>acceso y muestra un<br>"Permisos<br>**Alta**                     |

## **E2. Cronograma Estimado de Ejecución (Duración: 6 Semanas)**

Al ser un proyecto empresarial, el cronograma se divide en fases consecutivas utilizando un enfoque estructurado.

```
[Semana 1] Configuración del entorno de pruebas y diseño de casos detallados.
[Semana 2] Ejecución de Pruebas Funcionales (Módulos individuales: Ventas, RRHH).
[Semana 3] Ejecución de Pruebas Funcionales (Marketing, Postventa) y Pruebas de Roles.
[Semana 4] Pruebas de Integración de Datos (Flujos de extremo a extremo).
[Semana 5] Pruebas de Regresión y Corrección de Defectos (Bugs encontrados).
[Semana 6] Cierre de Pruebas, Informe Final y Firma de Aceptación (Sign-off).
```

## **Detalle de Horas Estimadas por Fase (Equipo de 2 QAs)**

- **Fase 1: Preparación (Semana 1):** 40 horas. Creación de datos base (empleados, productos, leads).

- **Fase 2: Ejecución Funcional e Integración (Semanas 2 a 4):** 120 horas. Pruebas manuales y automatización de flujos críticos.

- **Fase 3: Estabilización y Cierre (Semanas 5 y 6):** 80 horas. Re-testeo de errores corregidos por los desarrolladores.

## **Actividades:**

- Desarrollar un plan de un módulo/aplicación de un software académico o empresarial, como el ejemplo mostrado.

**NOTA** : Realizar un informe con carátula, desarrollar el caso propuesto, conclusiones y recomendaciones, y subir al classroom.

| **RÚBRICA:**                                                                            |                                                                                         |                                                                                         |                                                                                                   |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Inicio<br>00-10                                                                         | Proceso<br>11-13                                                                        | Logro previsto<br>14-17                                                                 | Logro satisfactorio<br>18-20                                                                      |
| Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio hasta un 50% | Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio hasta un 60% | Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio hasta un 80% | Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio<br>hasta<br>un<br>100% |
