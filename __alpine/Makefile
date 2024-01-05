all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./uploadNxRM.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.21.5 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/uploadNxRM.exe $(DESTDIR)$(PREFIX)/bin/uploadNxRM
	strip $(DESTDIR)$(PREFIX)/bin/uploadNxRM
