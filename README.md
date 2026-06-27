**Laboratorio Nº 06 E. P.:** Ingeniería de Sistemas (plan 2018)

**Asignatura:** IS489 Pruebas y Aseguramiento de calidad de Software

**Docente:** Mtro. Ing. Oscar Meneses Yaranga

**Semestre Académico** : 2026-I Presencial

**SESIÓN 0: Refactorización guiada por pruebas (TDD básico).**

## **I. OBJETIVOS**

Al término de esta experiencia, el estudiante será capaz de:

- Aplicar pruebas antes y después de refactorizar código.

- Realizar pruebas unitarias con JUnit 5 / pytest / Cobertura.

- Laboratorio 6 Código mejorado con pruebas y reporte comparativo.

## **II. EQUIPOS Y MATERIALES**

- Computadora personal (PC) y conexión a Internet.

- Compilador/intérprete de lenguajes de programación y/o IDE. ✓ Guía de laboratorio.

- Material bibliográfico con la información en páginas de la sesión de aprendizaje.

## **III. METODOLOGÍA Y ACTIVIDADES**

- Teoría de Desarrollo de Software (Pruebas de Software y Aseguramiento de calidad)

- Teoría de Sistemas de Información (Ingeniería del Software)

- Teoría de Proyectos de Desarrollo de Software

## **Revisión Teórica:**

## **Introducción al TDD**

**El desarrollo guiado por pruebas (TDD, por sus siglas en inglés)** es un enfoque de desarrollo de software que ha influido significativamente en la forma en que los desarrolladores escriben código, garantizando sistemas de software de mayor calidad y más fáciles de mantener. En esencia, TDD es tanto una filosofía como una metodología que enfatiza la importancia de escribir pruebas automatizadas antes de desarrollar la funcionalidad o el código propiamente dicho. Este enfoque invierte fundamentalmente el proceso tradicional de desarrollo de software, donde normalmente el código se escribe primero y se prueba después.

## **La filosofía detrás del TDD**

La filosofía del desarrollo guiado por pruebas (TDD) se basa en la idea de minimizar errores, mejorar la calidad del código y fomentar una comprensión más profunda de su propósito desde el principio. Anima a los desarrolladores a reflexionar sobre sus decisiones de diseño y requisitos antes de escribir el código, lo que da como resultado un código más claro, conciso y orientado a un propósito. Esta filosofía va más allá de las simples pruebas e influye en las estrategias generales de diseño y desarrollo de software, promoviendo una transición hacia prácticas de codificación más iterativas, receptivas y adaptativas.

## **Principios básicos del TDD**

TDD se basa en algunos principios fundamentales que guían a los desarrolladores a lo largo del proceso de desarrollo:

1. **Escribir una prueba fallida** : Antes de escribir código funcional, un desarrollador escribe una prueba automatizada que define una mejora deseada o una nueva función. Inicialmente, esta prueba fallará porque la función aún no se ha implementado.

2. **Superar la prueba** : El siguiente paso consiste en escribir la cantidad mínima de código necesaria para que la prueba se supere. Esto fomenta la simplicidad y la concentración en los requisitos.

3. **Refactorización del código** : Una vez superada la prueba, el código existente se refactoriza para cumplir con los estándares de limpieza, eficiencia y diseño. Esto puede implicar eliminar duplicaciones, mejorar la claridad o realizar otras mejoras que no alteren la funcionalidad.

## **El ciclo de TDD**

El proceso TDD es iterativo y consiste en ciclos repetidos de estos tres pasos (Rojo-VerdeRefactorización). Este ciclo fomenta la mejora continua del código y las pruebas, promoviendo una alta cobertura de código y asegurando que el software se desarrolle teniendo en cuenta las pruebas desde el principio.

## **Revertir el proceso de desarrollo tradicional**

Tradicionalmente, las pruebas son una fase posterior al desarrollo, lo que suele resultar en la identificación de errores y fallos de diseño al final del ciclo, cuya corrección puede ser costosa y llevar mucho tiempo. El desarrollo guiado por pruebas (TDD) invierte este proceso, priorizando las pruebas. Al escribir pruebas antes de escribir el código, los desarrolladores pueden garantizar que cada nueva funcionalidad se pruebe de inmediato, reduciendo errores y mejorando la calidad del diseño desde el principio. Este enfoque también garantiza que el código base siempre haya superado todas las pruebas, lo que aumenta la confianza en la fiabilidad y el comportamiento del software.

En esencia, TDD es más que una metodología de pruebas; es un enfoque integral para el desarrollo de software que prioriza la calidad, la documentación y la simplicidad. Al adoptar TDD, los desarrolladores pueden crear código más fiable, mantenible y libre de errores. Además, fomenta una mentalidad de mejora continua y adaptación, fundamental en el vertiginoso y cambiante panorama tecnológico actual.

## **¿Por qué es importante el TDD en el desarrollo de software?**

## • **Calidad de código mejorada**

Uno de los beneficios más inmediatos del TDD es la mejora significativa en la calidad del código. Al escribir pruebas antes del código propiamente dicho, los desarrolladores se ven obligados a analizar la funcionalidad y los casos límite desde el principio, lo que da como resultado implementaciones más reflexivas, robustas y resistentes a errores. Esta previsión ayuda a identificar posibles problemas al inicio del ciclo de desarrollo, cuando son más fáciles y menos costosos de solucionar.

## • **Diseño y arquitectura mejorados**

El desarrollo guiado por pruebas (TDD) anima a los desarrolladores a considerar el diseño de su código desde el principio. Dado que las pruebas se escriben primero, el código se diseña para que sea comprobable, lo que suele dar como resultado una arquitectura más limpia y modular. Este modularidad facilita el mantenimiento, la escalabilidad y la adaptación del software a los requisitos cambiantes a lo largo del tiempo.

## • **Facilitación de la refactorización**

La refactorización del código para mejorar su estructura, rendimiento o legibilidad sin modificar su comportamiento externo es una parte fundamental del desarrollo de software. El desarrollo guiado por pruebas (TDD) hace que la refactorización sea más segura y eficiente, ya que el conjunto de pruebas garantiza que la funcionalidad se mantenga intacta tras los cambios. Esta confianza permite a los desarrolladores mejorar continuamente el código base, manteniendo el software en óptimas condiciones y adaptable.

## • **Documentación y especificación**

Las pruebas escritas con la metodología TDD sirven como documentación viva del sistema. Describen claramente la función del código, lo cual es invaluable para los nuevos miembros del equipo o al revisar el código fuente después de un tiempo. Este aspecto de TDD garantiza que el conocimiento sobre la funcionalidad del sistema se conserve y sea fácilmente accesible.

## • **Reducción de la tasa de errores**

Al priorizar la cobertura de pruebas y las pruebas tempranas, el desarrollo guiado por pruebas (TDD) reduce significativamente la tasa de errores en los productos de software. Este enfoque proactivo para la identificación y resolución de errores da como resultado un software más fiable, aumenta la confianza del usuario y reduce el tiempo y los recursos dedicados a la depuración y corrección de problemas después del lanzamiento.

## • **Mejora la productividad del desarrollador**

Si bien el desarrollo guiado por pruebas (TDD) puede parecer que ralentiza el proceso de desarrollo inicialmente debido al tiempo adicional dedicado a escribir pruebas, en realidad mejora la productividad a largo plazo. La reducción de errores, las especificaciones más claras y un código base más limpio disminuyen el tiempo que los desarrolladores dedican a depurar y reelaborar código problemático. Además, la naturaleza iterativa del TDD, con sus ciclos de retroalimentación cortos, mantiene el ritmo de desarrollo, asegurando un progreso constante.

## • **Mejor colaboración en equipo**

El desarrollo guiado por pruebas (TDD) puede fomentar una mejor colaboración dentro de los equipos de desarrollo. Dado que las pruebas definen claramente la funcionalidad y las especificaciones de diseño, garantizan que todos los miembros del equipo tengan una comprensión común de los objetivos y el progreso del proyecto. Esta claridad facilita la coordinación de esfuerzos, especialmente en equipos grandes o distribuidos en diferentes ubicaciones.

En el ámbito del desarrollo de software, donde la complejidad y las exigencias de los proyectos no dejan de crecer, el TDD ofrece un enfoque estructurado y disciplinado que aborda muchos de los desafíos inherentes. Mejora la calidad del software, agiliza el proceso de desarrollo y se integra perfectamente con las metodologías de desarrollo ágil e iterativo. Al adoptar el TDD, los equipos no solo pueden mejorar su rendimiento inmediato, sino también garantizar la sostenibilidad y el éxito a largo plazo de sus proyectos de software.

## **IV. INDICACIONES**

Antes de empezar con el laboratorio, verificar tenerlos las aplicaciones de Office, puesto los documentos se redactarán en Word o similares, analizar los siguientes:

## **Comprender el ciclo TDD**

El ciclo TDD consta de tres etapas principales que los desarrolladores repiten para cada nueva característica o unidad de funcionalidad que desean implementar. Estas etapas están diseñadas para garantizar que el desarrollo de software se guíe por pruebas, lo que da como resultado un código de mayor calidad que cumple con los requisitos con precisión y es más fácil de mantener y adaptar con el tiempo.

## **1. Rojo: Escribe una prueba fallida**

- **Propósito** : El ciclo comienza con la fase "Roja", donde el desarrollador escribe una nueva prueba que define el comportamiento esperado de una característica o una pequeña funcionalidad que aún no existe. Esta prueba fallará de forma natural la primera vez que se ejecute, ya que la característica que prueba no se ha implementado. El fallo es intencional y sirve como punto de referencia para asegurar que los cambios de código posteriores marquen una diferencia tangible en la implementación de la funcionalidad deseada.

- **Implementación** : Para escribir una prueba fallida, es necesario comprender los requisitos y las especificaciones de la funcionalidad. La prueba debe ser concisa y centrarse en un único aspecto de la funcionalidad. Esta fase garantiza una definición clara de lo que significa el éxito para la funcionalidad en desarrollo.

## **2. Verde: Superar la prueba**

- **Objetivo** : La fase "verde" consiste en escribir la cantidad mínima de código necesaria para que la prueba fallida pase. Este paso no busca crear una solución perfecta al primer intento, sino lograr rápidamente un estado funcional que cumpla con los criterios de la prueba.

- **Implementación** : Aquí la clave está en la simplicidad y la eficacia. El desarrollador escribe el código justo y necesario para cumplir con los requisitos de la prueba, aunque la solución no sea la más elegante ni la más eficiente. El objetivo es superar la prueba y garantizar que la funcionalidad funcione según lo previsto.

## **3. Refactorizar: Mejorar el código**

- **Objetivo** : Una vez superada la prueba, el ciclo pasa a la fase de refactorización. En esta etapa, el desarrollador perfecciona el código, mejorando su estructura, legibilidad y rendimiento sin modificar su comportamiento externo. Este paso es fundamental para mantener la calidad y la facilidad de gestión del código a lo largo del tiempo.

- **Implementación** : La refactorización puede abarcar diversas actividades, desde renombrar variables para mayor claridad, reducir la duplicación y extraer métodos para un mejor modularidad, hasta aplicar patrones de diseño. Las pruebas escritas en la primera fase sirven como red de seguridad, garantizando que estas mejoras no alteren la funcionalidad.

## **La naturaleza iterativa del ciclo TDD**

El ciclo TDD es inherentemente iterativo. Tras completar un ciclo, el desarrollador inicia el siguiente escribiendo una nueva prueba para la siguiente funcionalidad o para otro aspecto de la aplicación. Este proceso iterativo fomenta el desarrollo gradual, donde las funcionalidades se construyen poco a poco, y cada paso se verifica mediante pruebas. Promueve un enfoque disciplinado del desarrollo de software, donde el progreso se valida continuamente y la calidad se integra desde el principio.

## **V. PROCEDIMIENTO**

Cumplir con las indicaciones dadas, para esto Ud. deberá seguir los siguientes pasos:

## **Desarrollo guiado por pruebas (TDD) en Java**

## **Puntos clave** :

- Comprender el flujo de trabajo de TDD: **Rojo** , **Verde** , **Refactorizar** .

- Primero, escribe la prueba que falla.

- Desarrollar código para que la prueba pase.

- Refactorización para obtener código limpio y optimizado.

- Ejemplos del uso de JUnit en Java.

## **1. El ciclo TDD: Rojo, Verde, Refactorización**

TDD sigue un ciclo simple y repetitivo:

- **Rojo** : Escribe una prueba para una nueva función, que fallará porque la función aún no está implementada.

- **Verde** : Escribe la cantidad mínima de código necesaria para que la prueba pase.

- **Refactorizar** : Limpiar el código, asegurándose de que siga las mejores prácticas sin que se rompa la prueba.

## **2. Escribe una prueba fallida (en rojo).**

El primer paso en TDD es escribir una prueba que defina una función o comportamiento que necesitas, pero que aún no se ha implementado.

Vamos a crear una clase de calculadora sencilla. Queremos escribir una prueba para verificar el método add().

`import org.junit.jupiter.api.Test; import static org.junit.jupiter.api.Assertions.assertEquals; public class CalculatorTest { @Test public void testAdd() { Calculator calculator = new Calculator(); int result = calculator.add(2, 3); assertEquals(5, result); } }`

En esta prueba, asumimos que existe un método add(int a, int b) que devuelve la suma de dos enteros. Sin embargo, como aún no hemos escrito la clase Calculator, la prueba fallará.

## **3. Desarrollar el código para pasar la prueba (Verde)**

Ahora, escribimos el código más sencillo necesario para que esta prueba pase.

`public class Calculator { public int add(int a, int b) { return a + b; } }`

En este punto, implementamos el método add() para que la prueba pase. Ejecuta la prueba de nuevo y ahora debería pasar.

## **4. Refactorizar (Limpiar el código)**

Una vez que la prueba se supera, el siguiente paso es revisar el código y refactorizarlo. En este caso, el código ya es limpio y sencillo, por lo que no se requiere una refactorización importante. Sin embargo, en casos más complejos, la refactorización ayuda a eliminar la duplicación y a mejorar la legibilidad.

## **Beneficios del TDD**

- **Confianza en el código** : Escribir pruebas primero garantiza que el código cumpla con el comportamiento esperado desde el principio.

- **Menos errores** : El desarrollo guiado por pruebas (TDD) ayuda a detectar errores en una etapa temprana, ya que se verifica cada funcionalidad antes de escribir su implementación.

- **Código más limpio** : La refactorización regular fomenta la escritura de código más limpio y eficiente.

- **Mejor diseño** : TDD te obliga a pensar en el diseño y la estructura de tu código antes de escribirlo.

## **Junit**

JUnit es un framework de pruebas unitarias para Java. Se utiliza para probar la unidad de código más pequeña, que es un método.

Aquí un ejemplo de una prueba JUnit:

`@Test public void testAdd() {`

`int result = calculator.add(1, 2); assertEquals(3, result); }`

Al realizar pruebas, debe usar la anotación @Test. El método _assertEquals_ se utiliza para comprobar si el resultado del método coincide con el resultado esperado. Si el resultado no coincide con el resultado esperado, la prueba fallará.

Aquí una prueba que fallará:

`@Test public void testAdd() { int result = calculator.add(1, 2); assertEquals(4, result); }`

Falla porque 1+2 no es igual a 4.

## **Mockito**

Mockito es un framework de simulación para Java. Se utiliza para simular las dependencias de la clase que se está probando.

Para usar Mockito, debes usar las anotaciones @Mock y @InjectMocks. La anotación @Mock se usa para crear un objeto simulado. La anotación @InjectMocks se usa para inyectar el objeto simulado en la clase que estás probando. Aquí tienes un ejemplo completo:

`@ExtendWith(MockitoExtension.class) public class UserServiceTest { @Mock private UserRepository userRepository; @InjectMocks private UserService userService; @Test public void testGetUserById() { User user = new User(); user.setId(1); user.setName("Beatriz"); when(userRepository.findById(1)).thenReturn(Optional.of(user)); User result = userService.getUserById(1); assertEquals(user, result); } }`

**En este ejemplo,** se está probando el método **getUserById** de la clase **UserService** . Se está simulando la clase **UserRepository** porque **UserService** depende de ella. Utiliza el método **when** para simular el método **findById** de **UserRepository** . El método **thenReturn** especifica qué devolverá **findById** ; en este caso, devolverá un objeto **Optional** que contiene un objeto **User. El método assertEquals** verifica si el resultado de **getUserById** coincide con el resultado esperado, que es el objeto **User** que creé manualmente.

## **Cobertura de pruebas**

La cobertura de pruebas es una medida de cuántas líneas de código se ejecutan durante las pruebas. Se utiliza para medir qué porcentaje del código se prueba. Existen diferentes tipos de cobertura de pruebas, pero la más común es la cobertura de líneas. Esta mide cuántas líneas de código se ejecutan durante las pruebas.

## **Creación de objetos simulados.**

Con Mockito, se puede crear fácilmente objetos simulados mediante el método `mock`. Estos objetos simulados se pueden configurar para que devuelvan valores o comportamientos específicos cuando se llamen a sus métodos.

`import org.mockito.Mockito; // Crea un objeto simulado UserService  userServiceMock = Mockito.mock(UserService.class); // Configura el objeto simulado para que devuelva un valor específico Mockito.when(userServiceMock.getUserById(1)).thenReturn(new  User (1, “John MY”));`

**La verificación** que ofrece Mockito permite comprobar si se han llamado métodos específicos de los objetos simulados con los argumentos esperados. Esto resulta muy útil para probar cómo interactúan las diferentes partes del código.

`// Verificar que se haya llamado a un método específico Mockito.verify(userServiceMock).getUserById( 1 );`

Mockito cuenta con numerosos mecanismos de comparación de argumentos que permiten definir reglas de coincidencia flexibles para los argumentos de los métodos. Esto resulta especialmente útil cuando se necesita verificar llamadas a métodos con argumentos dinámicos o complejos **.**

`// Verifica una llamada a un método con cualquier argumento Mockito.verify(userServiceMock).saveUser(Mockito.any(User.class));`

## **Implementación de un servicio de usuario.**

Vamos a mostrar el ciclo TDD creando una clase UserService simple usando JUnit y Mockito.

`import org.junit.jupiter.api.Test; import org.mockito.Mockito; import static org.junit.jupiter.api.Assertions.assertEquals; import static org.mockito.Mockito.when;`

`public class  UserServiceTest {`

`@Test public void  testGetUserById () {`

`// Arrange UserRepository userRepositoryMock  = Mockito.mock(UserRepository.class); when(userRepositoryMock.findById(1)).thenReturn(new  User (1 , "John MY" )); UserService userService  =  new  UserService (userRepositoryMock);`

`// Act User user  = userService.getUserById( 1`

`// Assert assertEquals( "John Doe" , user.getName()); } }`

## **Escribe el código mínimo de producción necesario para que la prueba pase:**

`public class  UserService {`

`private final UserRepository userRepository;`

`public  UserService (UserRepository userRepository) { this .userRepository = userRepository; } public User getUserById ( int id) { return userRepository.findById(id); } }`

## **Refactoriza el código asegurándote de que todas las pruebas pasen.**

Siguiendo el ciclo TDD con JUnit y Mockito, los desarrolladores pueden crear código robusto y bien probado paso a paso, mejorando la calidad del código y reduciendo el riesgo de errores.

## **Pruebas parametrizadas**

`import org.junit.jupiter.params.ParameterizedTest; import org.junit.jupiter.params.provider.ValueSource;`

`public class  CalculatorTest {`

`@ParameterizedTest`

`@ValueSource(ints = {1, 2, 3, 4, 5})`

`public void  testMultiplication ( int value)`

`Calculator calculator  =  new  Calculator (); int result  = calculator.multiply(value, 3 ); assertEquals(value * 3 , result);`

`} }`

## **Cómo usar JUnit y Mockito juntos**

1. Configura tu clase de prueba: crea una nueva clase para tus pruebas y usa la anotación **@Test** para marcar tus métodos de prueba.

2. Crear objetos simulados: Utilice Mockito para crear objetos simulados para las partes de su código que desea aislar.

3. Escribe tus pruebas: Escribe pruebas que utilicen los objetos simulados para simular diferentes escenarios y comprobar si tu código se comporta correctamente.

4. Ejecuta tus pruebas: Usa JUnit para ejecutar todas tus pruebas y ver los resultados.

## **Ejemplo**

Digamos que tienes una clase **Calculator** con un método **add** que suma dos números. Puedes escribir una prueba unitaria para este método usando JUnit y Mockito de la siguiente manera:

`import static org.junit.jupiter.api.Assertions.assertEquals; import org.junit.jupiter.api.Test; import org.mockito.Mockito;`

`class  CalculatorTest {`

`@Test void  testAdd () {`

`Calculator calculator  = Mockito.mock(Calculator.class);`

`Mockito.when(calculator.add( 2 , 3 )).thenReturn( 5`

`int result  = calculator.add( 2 , 3 );`

`assertEquals( 5 , result);`

`}`

`}`

## **VI. CASOS PROPUESTOS:**

## Revisar los siguientes:

Un sistema de procesamiento de pagos.

- Validar que los cálculos de impuestos se realicen correctamente.

- Comprobar que las restricciones de pago (montos mínimos, límites diarios) se respeten.

- Asegurar que los reembolsos se procesen según las políticas establecidas.

Un portal web de autoservicio.

- Confirmar que los botones, formularios y menús funcionan según lo esperado.

- Evaluar la navegación y los flujos de usuario.

- Verificar que los mensajes de error sean claros y útiles.

Un sistema de reservas en línea que consulta disponibilidad en diferentes proveedores.

- Verificar que las respuestas de la API se procesen correctamente.

- Evaluar cómo responde el sistema ante tiempos de espera o fallos en los servicios externos.

- Validar que los datos intercambiados sean consistentes.

Un e-commerce con alto volumen de transacciones.

- Simular compras con diferentes métodos de pago y validar resultados.

- Probar la correcta aplicación de descuentos y promociones.

- Detectar posibles errores en la generación de facturas.

## **Actividades:**

Tomando como ejemplo de proyecto de la lista, generar una implementación con TDD en cualquier lenguaje de programación moderna, mostrar las ejecuciones y subir a github.

**NOTA** : Hacer un informe con carátula y subir al classroom. Se puede poner capturas con detalles de cada código o comentarios legibles, agregar conclusiones y recomendaciones.

## **RÚBRICA:**

| **RÚBRICA:**                                                                            |                                                                                         |                                                                                         |                                                                                                   |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| Inicio<br>00-10                                                                         | Proceso<br>11-13                                                                        | Logro previsto<br>14-17                                                                 | Logro satisfactorio<br>18-20                                                                      |
| Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio hasta un 50% | Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio hasta un 60% | Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio hasta un 80% | Desarrolla<br>correctamente<br>las<br>prácticas<br>en<br>el<br>laboratorio<br>hasta<br>un<br>100% |

---

## **Bibliografía**

Jorgensen, P. Software Testing. CRC Press. 2014

Piattinni, G. M. Auditoría Informática. Un enfoque práctico. Alfaomega

Gilb, T. Software Inspection. Addison-Wesley Professional. 1994

https://www.browserstack.com/guide/tdd-in-java

https://stackoverflow.com/questions/31438065/how-to-do-test-driven-development-rightway

https://www.geeksforgeeks.org/software-testing/test-driven-development-using-junit5-andmockito/
