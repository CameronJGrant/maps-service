import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    distDir: "build",
    rewrites: async () => {
    	return [
     		{
       			source: "/v1/submissions/1/status/:statustype",
          		destination: "http://localhost:8081/v1/submissions/:statustype"
       		}
     	]
    }
};

export default nextConfig;
