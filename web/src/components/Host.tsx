import React from "react";
import styled from "styled-components";
import tw from "twin.macro";
import Highlight, { defaultProps } from "prism-react-renderer";
import duotoneDarkTheme from "prism-react-renderer/themes/dracula";
import nightOwlLightTheme from "prism-react-renderer/themes/nightOwlLight";

export const h1 = styled.h1(tw`font-black text-3xl mb-5`);
export const h2 = styled.h2(tw`font-bold text-2xl mb-2`);
export const h4 = styled.h4(tw`font-bold font-mono mb-2`);

export const p = styled.p(tw`mb-2`);

export const ul = styled.ul(tw`list-disc pl-5`);

export const code = styled.code(
  tw`bg-red-50 text-red-900 p-px px-1 text-sm border-red-100 border rounded inline-block`
);

export const pre = (props) => {
  const codeEl = props.children[0];
  return (
    <Highlight
      {...defaultProps}
      code={codeEl.props.children[0].replace(/\n$/, "")}
      language={
        codeEl.props.className
          ? codeEl.props.className.replace("language-", "")
          : undefined
      }
      theme={
        codeEl.props.className == "language-sh"
          ? duotoneDarkTheme
          : nightOwlLightTheme
      }
    >
      {({ className, style, tokens, getLineProps, getTokenProps }) => (
        <pre
          className={className}
          style={style}
          tw="border-black border-opacity-5 border rounded overflow-x-auto p-3 mb-2"
        >
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
