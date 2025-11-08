{
	pkgs ? import <nixpkgs> {},
	...
}:
pkgs.buildGoModule {
	pname = "fanbox2discord";
	version = "1.0.0";

	src = ./. ;

	doCheck = false;

	vendorHash = "sha256-CpEMYfg2MK8Bnm/hQWJ9ftYbzZBlWpVfc0rj2o4myws=";
}
