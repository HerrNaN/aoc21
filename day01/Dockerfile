FROM gradle:alpine as BASE
WORKDIR /app
COPY ./ .
RUN gradle build

FROM openjdk:17-alpine
COPY --from=BASE /app/build/libs/day01-1.0.0.jar /day01.jar
COPY input.txt .
CMD java -jar /day01.jar
