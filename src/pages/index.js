import * as React from "react";
import tw from "twin.macro";
import styled from "styled-components";
import heroBackground from "../images/james-brown-ascii-1.png";
import Features from "../components/features";
import { Helmet } from "react-helmet";

import installData from "../utils/data/install";
import usageData from "../utils/data/usage";

import colors from "../utils/colors";
import { textShadow, boxShadowSmall, boxShadowLarge } from "../utils/shadows";

const HeroSection = styled.section`
  ${tw`h-screen flex flex-col items-center justify-center`}

  background: url(${heroBackground}) no-repeat;
  background-size: cover;
`;

const HeroTitle = styled.h1`
  ${tw`text-6xl p-4 text-white`}

  background: ${colors.purple};
  ${textShadow};
  ${boxShadowSmall};
`;

const HeroBlurb = styled.p`
  ${tw`text-3xl p-4 text-white`}

  background: ${colors.orange};
  ${textShadow};
  ${boxShadowSmall};
`;

const SectionTitle = styled.h2`
  ${tw`text-5xl text-white mb-8`}
  ${textShadow}
`;

const Section = styled.section`
  ${tw`p-8`}
`;

const InstallationInstructions = styled(Section)`
  ${tw`flex flex-col items-center justify-center`}

  background: ${colors.purple};
`;

const InstallationTitle = styled.h2`
  ${tw`text-white`}
`;

const InstructionText = styled.p`
  ${tw`text-white`}
`;

const ExampleSection = styled(Section)`
  ${tw`flex flex-col items-center justify-center`}

  background: ${colors.orange};
`;

const ExampleUsage = styled.iframe`
  ${tw`bg-black`}

  min-height: 570px;
  min-width: 600px;
  border: none;
  ${boxShadowLarge};
`;

const Preformatted = styled.pre`
  ${tw`text-white px-4`}

  max-width: 600px;
  overflow-x: auto;
  background: ${colors.orange};
  ${boxShadowSmall};
  ${textShadow};
`;

const UsageSection = styled(Section)`
  ${tw`flex flex-col items-center justify-center`}

  background: ${colors.purple};

  p {
    ${tw`text-white p-4`};
    background: ${colors.orange};
    max-width: 500px;
    ${boxShadowSmall};
    ${textShadow};
  }
`;

const BlockTitle = styled.h3`
  ${tw`text-white text-left`}
  width: 500px;
  ${textShadow};
  max-width: 100%;
`;

const UsageCommandsList = styled.div`
  ${tw`text-white`}

  width: 500px;
  background: ${colors.orange};
  max-width: 100%;
  ${boxShadowSmall};
`;

const UsageCommandDetails = styled.div`
  ${tw`w-full flex flex-row justify-between`}
`;

const UsageCommandDetail = styled.div`
  ${tw`px-4 text-right`}
  width: 50%;

  &:last-child {
    ${tw`text-left relative`}

    &::before {
      ${tw`absolute`}
      content: "->";
      left: -6px;
    }
  }
`;

const UsageCommand = styled.div`
  ${tw`flex flex-col p-4 border-solid border-0 border-b border-white`}

  &::last-child {
    ${tw`border-none border-b-0`}
  }

  &.heading {
    ${tw`text-2xl`}
    ${UsageCommandDetail} {
      &:last-child {
        ${tw`text-left relative`}

        &::before {
          ${tw`absolute`}
          content: "->";
          left: -10px;
        }
      }
    }
  }
`;

const UsageCommandExamples = styled.div`
  ${tw`text-center w-full mt-2`}
`;

// markup
const IndexPage = () => {
  return (
    <>
      <Helmet>
        <title>git-machine | I feel good</title>
      </Helmet>
      <main>
        <HeroSection>
          <HeroTitle>git-machine</HeroTitle>
          <HeroBlurb>`just git on down` with the hippest flow around</HeroBlurb>
        </HeroSection>
        <ExampleSection>
          <SectionTitle>Example</SectionTitle>
          <ExampleUsage src="https://asciinema.org/a/400666/embed?size=small&speed=2" />
        </ExampleSection>
        <Features />
        <InstallationInstructions>
          <InstallationTitle>Installation Instructions</InstallationTitle>
          <InstructionText>
            This will create a .bin folder where the script was run. You will
            need to add that to your PATH.
          </InstructionText>
          <div>
            <Preformatted>
              <code>{installData.instructions}</code>
            </Preformatted>
          </div>
          <InstructionText>
            You can also pass in a directory already on your PATH.
          </InstructionText>
          <div>
            <Preformatted>
              <code>{installData.instructionsVerbose}</code>
            </Preformatted>
          </div>
        </InstallationInstructions>
        <UsageSection>
          <SectionTitle>Usage</SectionTitle>
          <p>
            The root executable is gitm. That is the only root command/alias
            enabled by default.
          </p>
          <p>
            Enable other aliases by using `gitm config aliases add`. This will
            use your config at `~/.git-machine/config` or it will use the
            default values of `just lets want need`.
          </p>
          <p>eg: "just get on down" instead of gitm on down</p>
          <BlockTitle>Command Recognition</BlockTitle>
          <p>
            git-machine uses pattern matching to determine the git command. If
            it recognizes a keyword, it drops insignificant text before it and
            proxies the rest of the arguments to git with the determined
            command.
          </p>
          <p>
            As an example `lets get on up` will recognize "up" and ignore "get"
            and "on".
          </p>
          <BlockTitle>Commands List</BlockTitle>

          <UsageCommandsList>
            <UsageCommand className="heading">
              <UsageCommandDetails>
                <UsageCommandDetail>gitm keyword</UsageCommandDetail>
                <UsageCommandDetail>git command</UsageCommandDetail>
              </UsageCommandDetails>
            </UsageCommand>
            {usageData.commands.map((command) => (
              <UsageCommand>
                <UsageCommandDetails>
                  <UsageCommandDetail>
                    {command.keywords.join(" | ")}
                  </UsageCommandDetail>
                  <UsageCommandDetail>{command.gitCommand}</UsageCommandDetail>
                </UsageCommandDetails>
                <UsageCommandExamples>
                  <div>
                    <b>example:</b>
                  </div>{" "}
                  {command.examples.join(" | ")}
                </UsageCommandExamples>
              </UsageCommand>
            ))}
          </UsageCommandsList>
        </UsageSection>
      </main>
    </>
  );
};

export default IndexPage;
