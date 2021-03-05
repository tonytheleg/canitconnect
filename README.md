# CanItConnect - a 100DaysOfCode Project

CanItConnect is a re-creation of a tool made at my previous employer to perform network tests from a container platform. CanItConnect can be deployed in platforms like PCF, Kubernetes, Docker Swarm, etc and allow testing connectivity to internal and external endpoints. Its purpose was to allow a simple API/UI for developers to test their apps (which would take the same network hops) have connectivity to their services without requiring all these network tools be available.

The tool was incredibly handy in troubleshooting firewall and connection issues, especially for those not strong in network testing tools on the command line. 

This project is a personal [100 Days of Code](https://www.100daysofcode.com/) project in an effort to break the tutorial death cycle and start having some fun wtih Go! Through this I'm hoping to learn a little more about Go web servers, creating both an API layer that can be used with tools like curl as well as a web console to provide a GUI test frontend as well. I'm hoping to dip into some WebAssembly as well!