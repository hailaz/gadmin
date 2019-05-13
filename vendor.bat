@echo off
echo will delete vendor dir...
rd /s vendor
go mod vendor
@pause