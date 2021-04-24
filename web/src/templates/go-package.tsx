import React, { FC } from "react";
import { PageProps, graphql, Link } from "gatsby";
import RehypeReact from "rehype-react";
import "twin.macro";

import Layout from "../components/Layout";
import * as components from "../components/Host";

const renderAst = new RehypeReact({
  createElement: React.createElement,
  components,
}).Compiler;

const GoPackage: FC<PageProps<any>> = ({ data }) => {
  const {
    htmlAst,
    parent: { name },
  } = data.markdownRemark;
  return (
    <Layout>
      <Link to="/" tw="text-xl my-6 inline-block">
        ← lukeharris.dev
      </Link>
      {renderAst(htmlAst)}
    </Layout>
  );
};

export const query = graphql`
  query($id: String!) {
    markdownRemark(id: { eq: $id }) {
      htmlAst
      parent {
        ... on File {
          name
        }
      }
    }
  }
`;

export default GoPackage;
