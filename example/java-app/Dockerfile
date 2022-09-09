FROM openjdk:11.0.6-jdk-slim as builder
COPY HelloArch.java /
RUN javac HelloArch.java

FROM openjdk:11-jre-slim
COPY --from=builder /HelloArch.class /
CMD ["java", "HelloArch"]