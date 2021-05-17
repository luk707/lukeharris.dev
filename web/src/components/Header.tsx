import React, { useState } from "react";
import styled from "styled-components";
import tw, { theme } from "twin.macro";
import { Link } from "gatsby";

const LOGO = `data:image/svg+xml;charset=UTF-8,%3csvg width='35' height='32' xmlns='http://www.w3.org/2000/svg'%3e%3cpath d='M13.37 25.2c.6 0 1.05-.14 1.36-.4.3-.28.46-.64.46-1.08 0-.43-.15-.79-.47-1.07-.31-.29-.76-.43-1.35-.43H6.03V9.2c0-.74-.17-1.3-.51-1.66A1.73 1.73 0 004.19 7c-.53 0-.97.18-1.3.55-.33.37-.49.91-.49 1.65v13.78c0 .78.17 1.35.52 1.7.34.35.91.52 1.7.52h8.75zm17.47.3c.54 0 .98-.18 1.32-.54.34-.37.51-.92.51-1.66V9.2c0-.74-.17-1.3-.5-1.66A1.73 1.73 0 0030.84 7c-.53 0-.97.18-1.3.55-.33.37-.5.91-.5 1.65v5.12h-7.82V9.2c0-.74-.17-1.3-.5-1.66A1.73 1.73 0 0019.39 7c-.55 0-.99.18-1.31.55-.32.37-.48.91-.48 1.65v14.1c0 .73.16 1.28.49 1.65.32.37.76.56 1.3.56s.98-.19 1.32-.55c.34-.37.51-.92.51-1.66v-6.02h7.83v6.02c0 .73.16 1.28.49 1.65.32.37.76.56 1.3.56z' fill='${theme`colors.gray.600`.replace(
  "#",
  "%23"
)}' fill-rule='nonzero'/%3e%3c/svg%3e`;

const CONTENT = [
  // {
  //   title: "Blog",
  //   to: "#blog",
  // },
  // {
  //   title: "Projects",
  //   to: "#projects",
  // },
  // {
  //   title: "Career",
  //   to: "#career",
  // },
];

const MobileMenu = styled.div(
  tw`absolute top-0 inset-x-0 p-2 transition transform origin-top-right sm:hidden`,
  ({ open }) =>
    open
      ? tw`opacity-100 scale-100`
      : tw`opacity-0 scale-95 pointer-events-none`
);

export const Header = () => {
  const [open, setOpen] = useState(false);
  return (
    <>
      <div tw="relative pt-6 mb-4">
        <nav
          tw="relative flex items-center justify-between sm:h-10 lg:justify-start"
          aria-label="Global"
        >
          <div tw="flex items-center flex-grow flex-shrink-0 lg:flex-grow-0">
            <div tw="flex items-center justify-between w-full md:w-auto">
              <a href="#">
                <span tw="sr-only">Workflow</span>
                <Link to="/">
                  <img tw="h-8 w-auto sm:h-10" src={LOGO} alt="LH" />
                </Link>
              </a>
              <div tw="-mr-2 flex items-center sm:hidden">
                <button
                  type="button"
                  tw="bg-white rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
                  aria-expanded="false"
                  onClick={() => setOpen(true)}
                >
                  <span tw="sr-only">Open main menu</span>
                  <svg
                    tw="h-6 w-6"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    aria-hidden="true"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M4 6h16M4 12h16M4 18h16"
                    />
                  </svg>
                </button>
              </div>
            </div>
          </div>

          <div tw="hidden sm:block sm:ml-10 sm:pr-4 sm:space-x-8">
            {CONTENT.map(({ title, to }) => (
              <a
                key={to}
                href={to}
                tw="font-medium text-gray-500 hover:text-gray-900"
              >
                {title}
              </a>
            ))}

            <a href="#" tw="font-medium text-indigo-600 hover:text-indigo-500">
              Log in
            </a>
          </div>
        </nav>
      </div>

      <MobileMenu open={open}>
        <div tw="rounded-lg shadow-md bg-white ring-1 ring-black ring-opacity-5 overflow-hidden">
          <div tw="px-5 pt-4 flex items-center justify-between">
            <div>
              <img tw="h-8 w-auto" src={LOGO} alt="LH" />
            </div>
            <div tw="-mr-2">
              <button
                type="button"
                tw="bg-white rounded-md p-2 inline-flex items-center justify-center text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-indigo-500"
                onClick={() => setOpen(false)}
              >
                <span tw="sr-only">Close main menu</span>
                <svg
                  tw="h-6 w-6"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12"
                  />
                </svg>
              </button>
            </div>
          </div>
          <div tw="px-2 pt-2 pb-3 space-y-1">
            {CONTENT.map(({ title, to }) => (
              <a
                key={to}
                href={to}
                tw="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:text-gray-900 hover:bg-gray-50"
              >
                {title}
              </a>
            ))}
          </div>
          <a
            href="#"
            tw="block w-full px-5 py-3 text-center font-medium text-indigo-600 bg-gray-50 hover:bg-gray-100"
          >
            Log in
          </a>
        </div>
      </MobileMenu>
    </>
  );
};
