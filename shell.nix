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
			dbmate
			(writeShellScriptBin "template" ''
				${dir}/template.sh "$@"
			'')
		];

		DATABASE_URL="sqlite:database.sqlite3";

	}
