cmake_minimum_required(VERSION 3.16)
project(example)

set(CMAKE_CXX_STANDARD 17)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

# CHANGE THIS TO THE APPROPRIATE DIRECTORY 
add_subdirectory(/deephaven/deephaven-core/cpp-client/deephaven deephaven_dir)

find_package(Arrow REQUIRED)
find_package(ArrowFlight REQUIRED HINTS ${Arrow_DIR})
find_package(Protobuf REQUIRED)
find_package(gRPC REQUIRED)
find_package(Threads REQUIRED)

add_executable(example main.cpp)

target_link_libraries(example deephaven::client)
