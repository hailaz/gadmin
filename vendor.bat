@echo off
echo will delete vendor dir...
rd /s vendor
echo govendor init
govendor init
echo govendor add +e
govendor add +e
echo finish
@pause