import React, { FC } from "react";
import { PageProps, graphql, Link } from "gatsby";
import RehypeReact from "rehype-react";
import "twin.macro";

import * as components from "../components/Host";
import { Header } from "../components/Header";
import Layout from "../components/Layout";

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
      <Header />
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
