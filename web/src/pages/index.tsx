import React, { FC } from "react";
import { PageProps, useStaticQuery, graphql, Link } from "gatsby";
import { Helmet } from "react-helmet";
import "twin.macro";

import { Header } from "../components/Header";
import { h2 as H2 } from "../components/Host";
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
      <Header />
      <H2>About me</H2>
      <ul tw="mb-4">
        <li>👨‍💻 Pronouns: He / Him</li>
        <li>🔭 I’m currently working on super secret work things</li>
        <li>🌱 I’m learning C/C++ &amp; Rust</li>
        <li>📫 Reach me on twitter @_lukeharris</li>
        <li>🎹 I mess around with audio synthesis in my spare time</li>
      </ul>
      <H2>Go packages</H2>
      <ul tw="list-disc pl-5">
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
              <Link
                to={`/${name}`}
                tw="bg-blue-50 border-b border-blue-300 text-blue-900 visited:bg-purple-50 visited:border-purple-300 visited:text-purple-900"
              >
                {name}
              </Link>
            </li>
          );
        })}
      </ul>
    </Layout>
  );
};

export default HomePage;
