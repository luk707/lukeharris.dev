import React, { FC } from "react";
import { PageProps, graphql } from "gatsby";

const GoPackage: FC<PageProps<any>> = ({ data }) => {
  const {
    goPackagesYaml: { import: importPath },
  } = data;
  return (
    <>
      <h1>lukeharris.dev/{importPath}</h1>
      <pre>{JSON.stringify(data, null, 4)}</pre>
    </>
  );
};

export const query = graphql`
  query($id: String!) {
    goPackagesYaml(id: { eq: $id }) {
      import
    }
  }
`;

export default GoPackage;
