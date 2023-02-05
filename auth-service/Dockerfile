FROM openjdk:11
ADD target/auth-service.jar auth-service.jar
ENTRYPOINT ["java", "-jar", "auth-service.jar"]
EXPOSE 8001