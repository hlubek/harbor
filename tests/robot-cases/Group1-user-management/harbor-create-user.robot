*** Settings ***
Documentation  It's an demo case to test robot and drone.
Resource  ../../resources/Util.robot
Default Tags  regression

*** Test Cases ***
Install Harbor to Test Server
    Log To Console  \nstart docker
    Run Keywords  Start Docker Daemon Locally
    Log To Console  \ndocker started success, config cfg
    Run Keywords  Config Harbor cfg
    Log To Console  \ncomplile and up harbor now
    Run Keywords  Compile and Up Harbor With Source Code
