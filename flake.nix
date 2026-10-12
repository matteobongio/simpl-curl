{
  description = "";
  inputs =
    {
      nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    };
  outputs = { self, nixpkgs, ... }@inputs:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs {
        inherit system;
        config.allowUnfree = true;
      };
    in
    {
      devShells.${system}.default =
        pkgs.mkShell
          {
            nativeBuildInputs = with pkgs; [
              go
              gopls
            ];
          };
       packages.${system}.default = pkgs.buildGoModule rec {
        name = "simpl-curl";

        src = ./.;

        vendorHash = null;
      };
    };
}

