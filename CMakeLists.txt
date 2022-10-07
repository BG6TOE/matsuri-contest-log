cmake_minimum_required(VERSION 3.0)
project(test_go)

set(TARGET mclguiserver_go)

file(GLOB_RECURSE SRCS *.go)
set(LIB libmclguiserver.a)

find_package(Git REQUIRED)
execute_process(COMMAND "${GIT_EXECUTABLE}" rev-parse HEAD OUTPUT_VARIABLE COMMIT_ID OUTPUT_STRIP_TRAILING_WHITESPACE)
project(ProjectName VERSION 1.0 DESCRIPTION "${COMMIT_ID}")

add_custom_command(OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/${LIB}
  DEPENDS ${SRCS}
  WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/shared
  COMMAND env GOPATH=${GOPATH} go build -buildmode=c-shared
  -o "${CMAKE_CURRENT_BINARY_DIR}/${LIB}"
  ${CMAKE_GO_FLAGS} ./...
  COMMENT "Building Go library")

add_custom_target(${TARGET} DEPENDS ${LIB} ${HEADER})
add_library(mclguiserver STATIC IMPORTED GLOBAL)
add_dependencies(mclguiserver ${TARGET})
set_target_properties(mclguiserver
  PROPERTIES
  IMPORTED_LOCATION ${CMAKE_CURRENT_BINARY_DIR}/${LIB}
  INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR}
  INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR})