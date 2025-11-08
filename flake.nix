{
  description = "Watches for new posts on Fanbox and show them on Discord";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }: {

    packages.x86_64-linux.fanbox2discord = nixpkgs.legacyPackages.x86_64-linux.callPackage ./package.nix {} ;

    packages.x86_64-linux.default = self.packages.x86_64-linux.fanbox2discord;

  };
}
