import {
  inputFieldStyle,
  dividerStyle,
  cardStyle,
  buttonStyle,
  gridStyle,
  cardTitleStyle,
  typographyStyle,
  colorSprinkle,
  buttonLinkStyle,
  messageStyle,
} from "./index.es.js";

const headings = document.querySelectorAll("h1");
for (const heading of headings) {
  heading.classList.add(...typographyStyle({ size: "headline-37" }).split(" "));
}

const inputs = document.querySelectorAll('input[type="text"]');
for (const e of inputs) {
  e.classList.add(
    inputFieldStyle,
    ...typographyStyle({ size: "small", type: "regular" }).split(" ")
  );
}

const textFieldLabels = document.querySelectorAll(".text-field-label");
for (const e of textFieldLabels) {
  e.classList.add(
    ...typographyStyle({ size: "small", type: "regular" }).split(" ")
  );
}

const dividers = document.querySelectorAll("hr");
console.log(dividerStyle());
for (const e of dividers) {
  e.classList.add(dividerStyle());
}

const cards = document.querySelectorAll(".card");
for (const e of cards) {
  e.classList.add(cardStyle());
}

const cardTitles = document.querySelectorAll(".card-title");
for (const e of cardTitles) {
  e.classList.add(cardTitleStyle);
}

const submitBtns = document.querySelectorAll('[type="submit"]');
for (const e of submitBtns) {
  e.classList.add(...buttonStyle({ size: "regular" }).split(" "));
}

const buttons = document.querySelectorAll('[type="button"]');
for (const e of buttons) {
  e.classList.add(...buttonStyle({ size: "regular" }).split(" "));
}

const grid32s = document.querySelectorAll(".grid-32");
for (const e of grid32s) {
  e.classList.add(...gridStyle({ gap: 32 }).split(" "));
}

const grid8s = document.querySelectorAll(".grid-8");
for (const e of grid8s) {
  e.classList.add(...gridStyle({ gap: 8 }).split(" "));
}

const typographyAlternatives = document.querySelectorAll(
  ".typography-alternative"
);
for (const e of typographyAlternatives) {
  e.classList.add(
    ...typographyStyle({ size: "caption", type: "regular" }).split(" "),
    ...colorSprinkle({ color: "foregroundMuted" }).split(" ")
  );
}

const buttonLinks = document.querySelectorAll(".button-link");
for (const e of buttonLinks) {
  e.classList.add(...buttonLinkStyle().split(" "));
}

const errorMessages = document.querySelectorAll(".error-message");
for (const e of errorMessages) {
  e.classList.add(...messageStyle({ severity: "error" }).split(" "));
}
