import * as React from "react";
import tw from "twin.macro";
import styled from "styled-components";
import heroBackground from "../images/james-brown-ascii-1.png";

import installData from "../utils/data/install";

import colors from "../utils/colors";

const HeroSection = styled.section`
  ${tw`h-screen flex flex-col items-center justify-center`}

  background: url(${heroBackground}) no-repeat;
  background-size: cover;
`;

const HeroTitle = styled.h1`
  ${tw`text-6xl p-4 text-white`}
  background: ${colors.purple};
  text-shadow: 1px 1px 0 black;
`;

const HeroBlurb = styled.p`
  ${tw`text-3xl p-4 text-white`}

  background: ${colors.orange};
  text-shadow: 1px 1px 0 black;
`;

const InstallationInstructions = styled.section`
  ${tw`flex items-center justify-center`}
`;

// markup
const IndexPage = () => {
  console.log({ installData });
  return (
    <main>
      <title>git-machine</title>
      <HeroSection>
        <HeroTitle>git-machine</HeroTitle>
        <HeroBlurb>`just git on down` with the hippest flow around</HeroBlurb>
      </HeroSection>
      <InstallationInstructions>
        <h2>Installation Instructions</h2>
        <pre>
          <code>{installData.instructions}</code>
        </pre>
      </InstallationInstructions>
    </main>
  );
};

export default IndexPage;
