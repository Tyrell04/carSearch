import { isDev } from "./helper";

class Http {
    async get(url: string, queryParams?: Record<string, string>, options?: RequestInit): Promise<any> {
        try {
            const response = await fetch(this.buildUrl(url, queryParams), options);
            if (!response.ok) {
                throw new Error(response.statusText);
            }
            return await response.json();
        } catch (error) {
            console.error(error);
            throw error;
        }
    }

    async post(url: string, body: any, options?: RequestInit): Promise<any> {
        try {
            const response = await fetch(url, {
                ...options,
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body),
            });
            if (!response.ok) {
                throw new Error(response.statusText);
            }
            return await response.json();
        } catch (error) {
            console.error(error);
            throw error;
        }
    }

    private buildUrl(url: string, queryParams?: Record<string, string>): string {
        if (!queryParams) {
            return url;
        }

        const searchParams = new URLSearchParams();
        Object.entries(queryParams).forEach(([key, value]) => {
            searchParams.append(key, value);
        });

        const baseUrl = isDev ? `http://localhost:8000` : '';
        return `${baseUrl}${url}?${searchParams.toString()}`;
    }
}

const request = new Http();

export function get(url: string, queryParams?: Record<string, string>, options?: RequestInit): Promise<any> {
    return request.get(url, queryParams, options);
}

export function post(url: string, body: any, options?: RequestInit): Promise<any> {
    return request.post(url, body, options);
}
