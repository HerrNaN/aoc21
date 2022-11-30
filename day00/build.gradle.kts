import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "1.6.0"
    application
}

group = "se.cygni.aoc"
version = "1.0.0"

repositories {
    mavenCentral()
    maven(
        url = uri("https://jitpack.io")
    )
}

dependencies {
    testImplementation(kotlin("test"))
    implementation("cc.ekblad.konbini:konbini:0.1.2")
}

tasks.test {
    useJUnitPlatform()
}

tasks.withType<KotlinCompile>() {
    kotlinOptions.jvmTarget = "17"
}

tasks.withType<Jar> {
    duplicatesStrategy = DuplicatesStrategy.WARN

    manifest {
        attributes["Main-Class"] = "AocTemplateKt"
    }
    configurations["compileClasspath"].forEach { file: File ->
        from(zipTree(file.absoluteFile))
    }
}

application {
    mainClass.set("AocTemplateKt")
}
