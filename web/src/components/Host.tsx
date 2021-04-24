import React from "react";
import styled from "styled-components";
import tw from "twin.macro";
import Highlight, { defaultProps } from "prism-react-renderer";
import theme from "prism-react-renderer/themes/github";

export const h1 = styled.h1(tw`font-black text-3xl mb-5`);
export const h2 = styled.h2(tw`font-bold text-2xl mb-2`);
export const h4 = styled.h4(tw`font-bold font-mono`);

export const p = styled.p(tw`mb-2`);

export const code = (props) => {
  console.log(props);
  return (
    <Highlight
      {...defaultProps}
      code={props.children[0].replace(/\n$/, "")}
      language={props.className.replace("language-", "")}
      theme={theme}
    >
      {({ className, style, tokens, getLineProps, getTokenProps }) => (
        <pre className={className} style={style} tw="rounded p-2 mb-2">
          {tokens.map((line, i) => (
            <div {...getLineProps({ line, key: i })}>
              {line.map((token, key) => (
                <span {...getTokenProps({ token, key })} />
              ))}
            </div>
          ))}
        </pre>
      )}
    </Highlight>
  );
};
