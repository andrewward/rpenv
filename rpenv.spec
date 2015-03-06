Name:   rpenv
Version:  1.0.1
Release:  1%{?dist}
Summary: displays env vars set from existing environment
Source0: rpenv
URL: https://github.com/rentpath/rpenv
BuildRoot:  %(mktemp -ud %{_tmppath}/%{name}-%{version}-%{release}-XXXXXX)
License:  Copyright

%description
displays env vars set from existing environment and loaded from config file in
specified environment (ci, qa, or prod) or executes command in the context of
the existing environment variables and ones loaded from a config file.

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}%{appdir}/
%{__install} -D -m 0655 %{SOURCE0} %{buildroot}%{_bindir}/%{name}


%clean
rm -rf %{buildroot}


%files
%defattr(-,root,root,-)
%{_bindir}/%{name}


%changelog
* Fri Feb 06 2015 Andrew Ward <award at rentpath dot com> 0.0.1
- Initial
