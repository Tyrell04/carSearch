class Http {
    get(url: string, queryParams?: Record<string, string>, options?: RequestInit): Promise<Response> {
        return fetch(this.buildUrl(url, queryParams), options);
    }

    post(url: string, body: any, options?: RequestInit): Promise<Response> {
        return fetch(url, {
            ...options,
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body),
        });
    }

    private buildUrl(url: string, queryParams?: Record<string, string>): string {
        if (!queryParams) {
            return url;
        }

        const searchParams = new URLSearchParams();
        Object.entries(queryParams).forEach(([key, value]) => {
            searchParams.append(key, value);
        });
        return `${url}?${searchParams.toString()}`;
    }
}

const request = new Http();

export function get(url: string, queryParams?: Record<string, string>, options?: RequestInit): Promise<Response> {
    return request.get(url, queryParams, options);
}

export function post(url: string, body: any, options?: RequestInit): Promise<Response> {
    return request.post(url, body, options);
}
