"use client";
import React from "react";
import { css } from "@emotion/css";

type Props = {
  children: React.ReactNode;
};

type ButtonStyleProps = {
  type: "primary" | "secondary";
};

const buttonStyle = ({ type }: ButtonStyleProps) => css`
  background-color: ${type === "primary" ? "blue" : "gray"};
  color: ${type === "primary" ? "white" : "black"};
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;

  &:hover {
    background-color: ${type === "primary" ? "darkblue" : "darkgray"};
  }
`;

export const Button: React.FC<Props> = ({ children }) => {
  return (
    <button className={buttonStyle({ type: "primary" })}>{children}</button>
  );
};
