import React, { FC } from "react";
import { PageProps, graphql, Link } from "gatsby";

const GoPackage: FC<PageProps<any>> = ({ data }) => {
  const {
    html,
    parent: { name },
  } = data.markdownRemark;
  return (
    <>
      <Link to="/">lukeharris.dev</Link>
      <div dangerouslySetInnerHTML={{ __html: html }} />
    </>
  );
};

export const query = graphql`
  query($id: String!) {
    markdownRemark(id: { eq: $id }) {
      html
      parent {
        ... on File {
          name
        }
      }
    }
  }
`;

export default GoPackage;
