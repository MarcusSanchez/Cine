import { z } from "zod";
import dotenv from "dotenv";

dotenv.config({ path: process.cwd() + "\\src\\env\\.env" });

const variables = z.object({
  API_URL: z.string().url(),
});

try {
  variables.parse(process.env);
} catch (e: any) {
  console.error("environment errors: ", e.errors);
  process.exit(1);
}

declare global {
  namespace NodeJS {
    interface ProcessEnv extends z.infer<typeof variables> {}
  }
}