all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./nxrmuploader.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.21.6 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/nxrmuploader.exe $(DESTDIR)$(PREFIX)/bin/nxrmuploader
	strip $(DESTDIR)$(PREFIX)/bin/nxrmuploader
