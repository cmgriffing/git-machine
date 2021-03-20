import * as React from "react";
import tw from "twin.macro";
import styled from "styled-components";
import colors from "../utils/colors";
import FeaturesImageSource from "../images/james-brown-ascii-3.png";

import { boxShadowLarge } from "../utils/shadows";

const FeaturesSection = styled.section`
  ${tw`text-gray-600`}
`;

const FeaturesContainer = styled.div`
  ${tw`container px-5 py-24 mx-auto flex flex-wrap`}
`;

const FeaturesImageContainer = styled.div`
  ${tw`lg:w-1/2 w-full mb-10 lg:mb-0 rounded-lg overflow-hidden`}

  border-radius: 8rem 2rem;
  ${boxShadowLarge}
`;

const FeaturesImage = styled.img`
  ${tw`object-cover object-center h-full w-full`}
`;

const FeaturesList = styled.div`
  ${tw`flex flex-col flex-wrap lg:py-6 -mb-10 lg:w-1/2 lg:pl-12 lg:text-left text-center`}
`;

const FeaturesListItem = styled.div`
  ${tw`flex flex-col mb-10 lg:items-start items-center`}
`;

const FeaturesListItemImageContainer = styled.div`
  ${tw`w-12 h-12 inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 mb-5`}
`;

const FeaturesListItemImage = styled.svg`
  ${tw`w-6 h-6`}
`;

const FeaturesListItemContentContainer = styled.div`
  ${tw`flex-grow`}
`;

const FeaturesListItemContentTitle = styled.div`
  ${tw`text-gray-900 text-lg font-medium mb-3`}
`;

const FeaturesListItemContentDescription = styled.div`
  ${tw`leading-relaxed text-base`}
`;

const features = [
  {
    title: "Permissive",
    description:
      "The commands just flow through you. The engine stays permissive to allow your energy its own space.",
    icon: "",
  },
  {
    title: "Conversational",
    description:
      "You are a human, damnit! Don't let a computer take that away from you.",
    icon: "",
  },
  {
    title: "Expressive",
    description: "Express yourself. Let your cli usage be as free as you are.",
    icon: "",
  },
];

// markup
const Features = () => {
  return (
    <FeaturesSection>
      <FeaturesContainer>
        <FeaturesImageContainer>
          <FeaturesImage alt="feature" src={FeaturesImageSource} />
        </FeaturesImageContainer>
        <FeaturesList>
          {features.map((feature) => (
            <FeaturesListItem>
              <FeaturesListItemImageContainer>
                <FeaturesListItemImage
                  fill="none"
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  viewBox="0 0 24 24"
                >
                  <path d="M22 12h-4l-3 9L9 3l-3 9H2"></path>
                </FeaturesListItemImage>
              </FeaturesListItemImageContainer>
              <FeaturesListItemContentContainer>
                <FeaturesListItemContentTitle>
                  {feature.title}
                </FeaturesListItemContentTitle>
                <FeaturesListItemContentDescription>
                  {feature.description}
                </FeaturesListItemContentDescription>
              </FeaturesListItemContentContainer>
            </FeaturesListItem>
          ))}
        </FeaturesList>
      </FeaturesContainer>
    </FeaturesSection>
  );
};

export default Features;
