Name:		   ms-agent
Version:	   %{version}
Release: 	1%{?alphatag:.%{alphatag}}%{?dist}
Summary: 	A tools send zabbix alerts to ZbxTable
Group:		Applications/Internet
License:	   Apache-2.0
URL:		   https://zbxtable.cactifans.com
Source0:	   ms-agent-%{version}%{?alphatag:%{alphatag}}.tar.gz

Buildroot:	%{_tmppath}/ms-agent-%{version}-%{release}-root-%(%{__id_u} -n)

%description
A tools send zabbix alerts to ZbxTable

%global debug_package %{nil}

%prep
%setup0 -q -n ms-agent-%{version}%{?alphatag:%{alphatag}}

%build

%install

rm -rf $RPM_BUILD_ROOT

# install necessary directories
mkdir -p $RPM_BUILD_ROOT%{_sysconfdir}/ms-agent
mkdir -p $RPM_BUILD_ROOT%{_prefix}/lib/zabbix/alertscripts

# install  binaries and conf
install -m 0755 -p ms-agent $RPM_BUILD_ROOT%{_prefix}/lib/zabbix/alertscripts/
install -m 0755 -p app.ini $RPM_BUILD_ROOT%{_sysconfdir}/ms-agent/

exit 0

%clean
rm -rf $RPM_BUILD_ROOT


%define __debug_install_post   \
   %{_rpmconfigdir}/find-debuginfo.sh %{?_find_debuginfo_opts} "%{_builddir}/%{?buildsubdir}"\
%{nil}

%files
%defattr(755,root,root,755)
%dir %{_prefix}/lib/zabbix/alertscripts/
%dir %{_sysconfdir}/ms-agent/
%{_prefix}/lib/zabbix/alertscripts/ms-agent
%{_sysconfdir}/ms-agent/app.ini
