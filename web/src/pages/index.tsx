import React, { FC } from "react";
import { PageProps, useStaticQuery, graphql, Link } from "gatsby";
import { Helmet } from "react-helmet";

import Layout from "../components/Layout";

const HomePage: FC<PageProps> = ({ path }) => {
  const data = useStaticQuery(graphql`
    query {
      allMarkdownRemark {
        edges {
          node {
            id
            parent {
              ... on File {
                name
                relativePath
              }
            }
          }
        }
      }
    }
  `);
  return (
    <Layout>
      <Helmet>
        <meta
          name="go-import"
          content="lukeharris.dev git https://github.com/luk707/lukeharris.dev"
        />
      </Helmet>
      <h1>lukeharris.dev</h1>
      <h2>Go packages</h2>
      <ul>
        {data.allMarkdownRemark.edges.map(({ node }) => {
          const {
            id,
            parent: { name, relativePath },
          } = node;
          const [type] = relativePath.split("/");
          if (type !== "go-packages") {
            return null;
          }
          return (
            <li key={id}>
              <Link to={`/${name}`}>{name}</Link>
            </li>
          );
        })}
      </ul>
    </Layout>
  );
};

export default HomePage;
