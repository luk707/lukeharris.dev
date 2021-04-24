import React, { FC } from "react";
import { PageProps, graphql } from "gatsby";
import { Helmet } from "react-helmet";

const GoPackage: FC<PageProps<any>> = ({ data }) => {
  const {
    goPackagesYaml: { import: importPath },
  } = data;
  return (
    <>
      <Helmet>
        <meta
          name="go-import"
          content={`lukeharris.dev/${importPath} git https://github.com/luk707/lukeharris.dev/${importPath}`}
        />
      </Helmet>
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
