import { parse } from 'path';
import { readFileSync } from 'fs';
import { join } from 'path';

export default class Parser {
  private filePath: string;

  constructor(filePath: string) {
    this.filePath = filePath;
  }

  public parseFile(): { [key: string]: any } {
    const data = readFileSync(this.filePath, 'utf8');
    const json = JSON.parse(data);
    return json;
  }

  public parsePath(): { [key: string]: any } {
    const path = parse(this.filePath);
    return path;
  }

  public getFileName(): string {
    return parse(this.filePath).name;
  }

  public getDirectoryPath(): string {
    return join(parse(this.filePath).dir, '');
  }
}