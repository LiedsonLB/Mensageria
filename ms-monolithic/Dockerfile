# Etapa de build com Maven
FROM maven:3.9.6-eclipse-temurin-17 AS builder

WORKDIR /app

# Etapa 1: Cache de dependências
COPY pom.xml ./
RUN mvn dependency:go-offline

# Etapa 2: Copia todo o projeto e gera o .jar
COPY . .
RUN mvn clean package -DskipTests

# Etapa final: imagem leve apenas com o JAR
FROM eclipse-temurin:17-jre-alpine

WORKDIR /app

COPY --from=builder /app/target/*.jar app.jar

EXPOSE 8081

CMD ["java", "-jar", "app.jar"]