# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(gui_lib_library VERSION 0.0.1 LANGUAGES C)

set(TARGET gui_lib_go)

file(GLOB_RECURSE SRCS ${CMAKE_CURRENT_SOURCE_DIR}/go_lib/*.go)

if (UNIX)
set(LIB libgui_lib.so)
endif (UNIX)

if (WIN32)
set(LIB gui_lib.dll)
endif (WIN32)

execute_process(COMMAND git rev-parse HEAD OUTPUT_VARIABLE COMMITID OUTPUT_STRIP_TRAILING_WHITESPACE)
execute_process(COMMAND git describe --tags OUTPUT_VARIABLE VERSION OUTPUT_STRIP_TRAILING_WHITESPACE)
project(ProjectName VERSION 1.0 DESCRIPTION "${COMMITID}")

add_custom_command(OUTPUT "${CMAKE_CURRENT_BINARY_DIR}/${LIB}"
  DEPENDS ${SRCS}
  WORKING_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}"
  COMMAND go build -buildmode=c-shared -o "${CMAKE_CURRENT_BINARY_DIR}/${LIB}" -ldflags "-X version.GitCommit=${COMMITID} -X version.Version=${VERSION}"
  ${CMAKE_GO_FLAGS} ${CMAKE_CURRENT_SOURCE_DIR}/go_lib/shared/shared.go
  COMMENT "Building Go library")

add_custom_target(${TARGET} ALL DEPENDS "${CMAKE_CURRENT_BINARY_DIR}/${LIB}" ${HEADER})
add_library(gui_lib SHARED IMPORTED GLOBAL)
add_dependencies(gui_lib ${TARGET})
set_target_properties(gui_lib
  PROPERTIES
  IMPORTED_LOCATION "${CMAKE_CURRENT_BINARY_DIR}/${LIB}"
  INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR}
  OUTPUT_NAME "${CMAKE_CURRENT_BINARY_DIR}/${LIB}"
  INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR})

set_target_properties(${TARGET} 
  PROPERTIES
  PUBLIC_HEADER mcl.h
  OUTPUT_NAME "${CMAKE_CURRENT_BINARY_DIR}/${LIB}")
