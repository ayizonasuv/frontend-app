// parser.js
import { parse } from 'node-html-parser';

const parseHtml = (htmlString) => {
  try {
    const parser = new parse(htmlString);
    return parser;
  } catch (error) {
    console.error('Error parsing HTML:', error);
    return null;
  }
};

export default parseHtml;