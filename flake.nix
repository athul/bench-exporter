{
  description = "Bench Exporter";
  inputs = { nixpkgs.url = "github:nixos/nixpkgs"; };
  outputs = { self, nixpkgs }:
  let pkgs = nixpkgs.legacyPackages.x86_64-linux;
  in {
    devShell.x86_64-linux =
      pkgs.mkShell {
        buildInputs = [
          pkgs.bash
          pkgs.go
          pkgs.gotools
          pkgs.air
          pkgs.gopls
        ];
      };
  };
}
