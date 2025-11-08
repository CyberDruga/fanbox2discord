{
pkgs ? import <nixpkgs> {},
...
}:
let
	dir = toString ./. ;
in
	with pkgs; mkShell {
		packages = [
			sqlc
			sqlite
			(writeShellScriptBin "template" ''
				${dir}/template.sh "$@"
			'')
		];

	}
