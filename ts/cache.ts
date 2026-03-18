import { iPagination, iRelease } from "./global.d";
import * as fs from "fs";

const getCollectionCache = async () => {
  // open ./cache/collection.json if exists
  try {
    return JSON.parse(fs.readFileSync("./cache/collection.json", "utf-8"));
  } catch {
    return [];
  }
};

const setCollectionCache = async (collection: iRelease[]) => {
  fs.writeFileSync("./cache/collection.json", JSON.stringify(collection));
};

const isCollectionCached = async (
  pagination: iPagination
): Promise<boolean> => {
  const cache = await getCollectionCache();
  return cache.length === pagination.items;
};

export { getCollectionCache, setCollectionCache, isCollectionCached };

export default { getCollectionCache, setCollectionCache, isCollectionCached };
