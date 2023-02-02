cmake_minimum_required(VERSION 3.23)
project(go_library_sample)
set(CMAKE_CXX_STANDARD 11)

# Variables
set(GO_STATIC_LIBRARY_NAME go_library.a)

# Dependencies
find_program(GOLANG_BINARY "go")
IF(NOT GOLANG_BINARY)
    message(FATAL_ERROR "Cannot find the \"go\" binary, please install it and add it to PATH")
ENDIF()
add_custom_target(GO_STATIC_LIBRARY_TARGET
    DEPENDS "${PROJECT_BINARY_DIR}/go-library/${GO_STATIC_LIBRARY_NAME}")
add_custom_command(OUTPUT "${PROJECT_BINARY_DIR}/go-library/${GO_STATIC_LIBRARY_NAME}"
    COMMAND ${GOLANG_BINARY} build -o "${PROJECT_BINARY_DIR}/go-library/${GO_STATIC_LIBRARY_NAME}" -buildmode=c-archive .
    WORKING_DIRECTORY "${PROJECT_SOURCE_DIR}/go-library"
    DEPENDS "${PROJECT_SOURCE_DIR}/go-library")
add_library(GO_STATIC_LIBRARY STATIC IMPORTED)
set_target_properties(GO_STATIC_LIBRARY PROPERTIES IMPORTED_LOCATION ${PROJECT_BINARY_DIR}/go-library/${GO_STATIC_LIBRARY_NAME} LINKER_LANGUAGE CXX)
add_dependencies(GO_STATIC_LIBRARY GO_STATIC_LIBRARY_TARGET)

# Project
include_directories(include src "${PROJECT_BINARY_DIR}/go-library")
file(GLOB_RECURSE ALL_SOURCES
    "${PROJECT_SOURCE_DIR}/src/*.cpp"
    "${PROJECT_SOURCE_DIR}/src/*.c")
add_executable(${PROJECT_NAME} ${ALL_SOURCES})
target_link_libraries(${PROJECT_NAME} GO_STATIC_LIBRARY)
