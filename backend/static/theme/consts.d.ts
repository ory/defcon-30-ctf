export declare type Font = {
    fontFamily: string;
    fontStyle: string;
};
export declare type BreakPoints = {
    sm: string;
    md: string;
    lg: string;
    xl: string;
    xl2: string;
};
export declare const defaultBreakpoints: BreakPoints;
export declare type Theme = {
    accent: {
        def: string;
        muted: string;
        emphasis: string;
        disabled: string;
        subtle: string;
    };
    foreground: {
        def: string;
        muted: string;
        subtle: string;
        disabled: string;
        onDark: string;
        onAccent: string;
        onDisabled: string;
    };
    background: {
        surface: string;
        canvas: string;
    };
    error: {
        def: string;
        subtle: string;
        muted: string;
        emphasis: string;
    };
    success: {
        emphasis: string;
    };
    border: {
        def: string;
    };
    text: {
        def: string;
        disabled: string;
    };
    input: {
        background: string;
        disabled: string;
        placeholder: string;
        text: string;
    };
} & Font;
export declare const defaultFont: {
    fontFamily: string;
    fontStyle: string;
};
export declare const defaultLightTheme: Theme;
export declare const defaultDarkTheme: Theme;
//# sourceMappingURL=consts.d.ts.map