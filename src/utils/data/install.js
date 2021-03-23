const data = {
  instructions: `
curl -o- -s https://raw.githubusercontent.com/cmgriffing/git-machine/main/install.sh | bash
  `,
  instructionsVerbose: `
curl -o- -s https://raw.githubusercontent.com/cmgriffing/git-machine/main/install.sh | bash -s -- -b /usr/local/bin/
  `,
};

export default data;
