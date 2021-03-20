const usageData = {
  commands: [
    {
      keywords: ["about"],
      gitCommand: "branch",
      details: "",
      examples: ["lets talk about some-feature"],
    },

    {
      keywords: ["into"],
      gitCommand: "checkout",
      details: "",
      examples: ["lets get into -b some-feature"],
    },

    {
      keywords: ["say", "shout", "yell"],
      gitCommand: "commit",
      details: "",
      examples: ['just say -m "fix: the system"'],
    },

    {
      keywords: ["me", "yourself"],
      gitCommand: "blame",
      details: "",
      examples: ["just take a look at yourself", "lets admit it was me"],
    },

    {
      keywords: ["minute", "scene"],
      gitCommand: "stash",
      details: "",
      examples: ["just stay on the scene", "lets take a minute"],
    },

    {
      keywords: ["started", "begin"],
      gitCommand: "init",
      details: "",
      examples: ["lets get this party started", "just begin"],
    },

    {
      keywords: ["there"],
      gitCommand: "remote",
      details: "",
      examples: [
        "just look over there add origin main",
        "lets see whats in there -v",
      ],
    },

    {
      keywords: ["out", "lost"],
      gitCommand: "rm",
      details: "",
      examples: ["just get lost man.md"],
    },

    {
      keywords: ["together"],
      gitCommand: "merge",
      details: "",
      examples: ["lets get together main"],
    },

    {
      keywords: ["yeah", "woo"],
      gitCommand: "status",
      details: "",
      examples: ["lets hear ya say yeah", "just woo"],
    },

    {
      keywords: ["stuff", "mess"],
      gitCommand: "diff",
      details: "",
      examples: ["lets check this stuff", "just look at this mess"],
    },

    {
      keywords: ["break ... down"], // lift...up and build...up
      gitCommand: "rebase",
      details:
        'This "down" will take precedence over the (pull) command if accompanied by break.',
      examples: ["just break it on down"],
    },

    {
      keywords: ["up"],
      gitCommand: "push",
      details: "",
      examples: ["just get on up"],
    },

    {
      keywords: ["down"],
      gitCommand: "pull",
      details: "",
      examples: ["lets get this party started", "just begin"],
    },

    {
      keywords: ["the", "that"],
      gitCommand: "add",
      details: "",
      examples: ["lets get the index.html", "just handle that index.js"],
    },
  ],
};

export default usageData;
