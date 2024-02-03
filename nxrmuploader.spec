%define debug_package   %{nil}
%define _build_id_links none
%define _name nxrmuploader
%define _prefix /opt
%define _version 1.50.00
%define _rel 0
%define _arch x86_64
%define _binaryname nxrmuploader

Name:       nxrmuploader
Version:    %{_version}
Release:    %{_rel}
Summary:    NxRM binary package uploader

Group:      DevOps tools
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/devops/nxrmuploader

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
NxRM binary package uploader

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
if getent group devops > /dev/null; then
  exit 0
else
  if getent group 2500; then
    groupadd devops
  else
    groupadd -g 2500 devops
  fi
fi
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post
cd /opt/bin
sudo chgrp -R devops .
sudo chmod 775 /opt/bin/uploadNxRM

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Fri Feb 02 2024 RPM Builder <builder@famillegratton.net> 1.50.00-0
- minor refactoring, removal of un-needed files (jean-
  francois@famillegratton.net)
- Rewrote uploadFile(). version bump (jean-francois@famillegratton.net)
- Make struct members exportable (jean-francois@famillegratton.net)
- Getting ready to rewrite uploadFile() (jean-francois@famillegratton.net)
- Sync before branching out (jean-francois@famillegratton.net)
- Fix arch issue when fetching a new GOLANG version
  (builder@famillegratton.net)

* Fri Jan 05 2024 RPM Builder <builder@famillegratton.net> 1.01.00-0
- Doc update to reflect the latest changes (jean-francois@famillegratton.net)
- Completed passwd encryption and env add (jean-francois@famillegratton.net)
- Version bump (jean-francois@famillegratton.net)
- APK build fix (builder@famillegratton.net)
- Fixed perms on multiple files (builder@famillegratton.net)

* Wed Jan 03 2024 RPM Builder <builder@famillegratton.net> 1.00.00-0
- new package built with tito

