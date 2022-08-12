import { RecipeVariants } from '@vanilla-extract/recipes';
export declare const buttonSocialStyle: import("@vanilla-extract/recipes/dist/declarations/src/types").RuntimeFn<{
    size: {
        small: {
            fontSize: string;
            lineHeight: string;
        };
        medium: {
            fontSize: string;
            lineHeight: string;
        };
        large: {
            fontSize: string;
            lineHeight: string;
            padding: string;
        };
    };
    variant: {
        regular: {
            fontWeight: number;
            fontStyle: "normal";
        };
        semibold: {
            fontWeight: number;
            fontStyle: "normal";
        };
    };
}>;
export declare const buttonSocialTitleStyle: string;
export declare const buttonSocialIconStyle: import("@vanilla-extract/recipes/dist/declarations/src/types").RuntimeFn<{
    size: {
        small: {};
        medium: {};
        large: {
            paddingRight: string;
            fontSize: string;
        };
    };
}>;
export declare type ButtonSocialStyle = RecipeVariants<typeof buttonSocialStyle>;
//# sourceMappingURL=button-social.css.d.ts.map