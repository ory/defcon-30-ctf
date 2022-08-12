const pxToRem = (...px) => px.map((x) => `${x / 16}rem`).join(" ");
const defaultBreakpoints = {
  sm: pxToRem(640),
  md: pxToRem(768),
  lg: pxToRem(1024),
  xl: pxToRem(1280),
  xl2: pxToRem(1536)
};
const defaultFont = {
  fontFamily: "Inter",
  fontStyle: "normal"
};
({
  fontFamily: "Inter",
  textDecoration: "none",
  fontSize: pxToRem(16),
  lineHeight: pxToRem(28)
});
const defaultLightTheme = {
  ...defaultFont,
  accent: {
    def: "#3D53F5",
    muted: "#6475F7",
    emphasis: "#3142C4",
    disabled: "#E0E0E0",
    subtle: "#eceefe"
  },
  foreground: {
    def: "#171717",
    muted: "#616161",
    subtle: "#9E9E9E",
    disabled: "#BDBDBD",
    onDark: "#FFFFFF",
    onAccent: "#FFFFFF",
    onDisabled: "#e0e0e0"
  },
  background: {
    surface: "#FFFFFF",
    canvas: "#FCFCFC"
  },
  error: {
    def: "#9c0f2e",
    subtle: "#fce8ec",
    muted: "#e95c7b",
    emphasis: "#DF1642"
  },
  success: {
    emphasis: "#18A957"
  },
  border: {
    def: "#E0E0E0"
  },
  text: {
    def: "#FFFFFF",
    disabled: "#757575"
  },
  input: {
    background: "#FFFFFF",
    disabled: "#E0E0E0",
    placeholder: "#9E9E9E",
    text: "#424242"
  }
};
const defaultDarkTheme = {
  ...defaultFont,
  accent: {
    def: "#6475f7",
    disabled: "#757575",
    muted: "#3142c4",
    emphasis: "#3d53f5",
    subtle: "#0c1131"
  },
  foreground: {
    def: "#FFFFFF",
    muted: "#ddd9f7",
    subtle: "#9a8ce8",
    onDark: "#FFFFFF",
    onAccent: "#FFFFFF",
    onDisabled: "#e0e0e0",
    disabled: "#bdbdbd"
  },
  background: {
    surface: "#110d2b",
    canvas: "#090616"
  },
  border: {
    def: "#221956"
  },
  error: {
    def: "#e95c7b",
    subtle: "#2d040d",
    muted: "#9c0f2e",
    emphasis: "#df1642"
  },
  success: {
    emphasis: "#18a957"
  },
  input: {
    background: "#FFFFFF",
    text: "#424242",
    placeholder: "#9e9e9e",
    disabled: "#eeeeee"
  },
  text: {
    def: "#FFFFFF",
    disabled: "#757575"
  }
};
var theme_css_ts_vanilla = "";
var card_css_ts_vanilla = "";
function _defineProperty$1(obj, key, value) {
  if (key in obj) {
    Object.defineProperty(obj, key, {
      value,
      enumerable: true,
      configurable: true,
      writable: true
    });
  } else {
    obj[key] = value;
  }
  return obj;
}
function ownKeys$1(object, enumerableOnly) {
  var keys = Object.keys(object);
  if (Object.getOwnPropertySymbols) {
    var symbols = Object.getOwnPropertySymbols(object);
    if (enumerableOnly) {
      symbols = symbols.filter(function(sym) {
        return Object.getOwnPropertyDescriptor(object, sym).enumerable;
      });
    }
    keys.push.apply(keys, symbols);
  }
  return keys;
}
function _objectSpread2$1(target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i] != null ? arguments[i] : {};
    if (i % 2) {
      ownKeys$1(Object(source), true).forEach(function(key) {
        _defineProperty$1(target, key, source[key]);
      });
    } else if (Object.getOwnPropertyDescriptors) {
      Object.defineProperties(target, Object.getOwnPropertyDescriptors(source));
    } else {
      ownKeys$1(Object(source)).forEach(function(key) {
        Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key));
      });
    }
  }
  return target;
}
var shouldApplyCompound = (compoundCheck, selections, defaultVariants) => {
  for (var key of Object.keys(compoundCheck)) {
    var _selections$key;
    if (compoundCheck[key] !== ((_selections$key = selections[key]) !== null && _selections$key !== void 0 ? _selections$key : defaultVariants[key])) {
      return false;
    }
  }
  return true;
};
var createRuntimeFn = (config) => (options) => {
  var className = config.defaultClassName;
  var selections = _objectSpread2$1(_objectSpread2$1({}, config.defaultVariants), options);
  for (var variantName in selections) {
    var _selections$variantNa;
    var variantSelection = (_selections$variantNa = selections[variantName]) !== null && _selections$variantNa !== void 0 ? _selections$variantNa : config.defaultVariants[variantName];
    if (variantSelection != null) {
      var selection = variantSelection;
      if (typeof selection === "boolean") {
        selection = selection === true ? "true" : "false";
      }
      var selectionClassName = config.variantClassNames[variantName][selection];
      if (selectionClassName) {
        className += " " + selectionClassName;
      }
    }
  }
  for (var [compoundCheck, compoundClassName] of config.compoundVariants) {
    if (shouldApplyCompound(compoundCheck, selections, config.defaultVariants)) {
      className += " " + compoundClassName;
    }
  }
  return className;
};
var cardStyle = createRuntimeFn({ defaultClassName: "_660nzl0", variantClassNames: {}, defaultVariants: {}, compoundVariants: [] });
var cardTitleStyle = "_660nzl1";
var message_css_ts_vanilla = "";
var messageStyle = createRuntimeFn({ defaultClassName: "_1qj4dn90", variantClassNames: { severity: { error: "_1qj4dn91", success: "_1qj4dn92", disabled: "_1qj4dn93" } }, defaultVariants: {}, compoundVariants: [] });
var divider_css_ts_vanilla = "";
var dividerStyle = createRuntimeFn({ defaultClassName: "_3ldkmt0", variantClassNames: { sizes: { fullWidth: "_3ldkmt1" } }, defaultVariants: {}, compoundVariants: [] });
var typography_css_ts_vanilla = "";
var inputTypographyStyle = createRuntimeFn({ defaultClassName: "_1j36fun0", variantClassNames: { size: { "16": "_1j36fun1", "18": "_1j36fun2" }, type: { regular: "_1j36fun3", semiBold: "_1j36fun4" } }, defaultVariants: { size: 16, type: "regular" }, compoundVariants: [] });
var typographyStyle = createRuntimeFn({ defaultClassName: "_1j36fun5", variantClassNames: { size: { tiny: "_1j36fun6", xsmall: "_1j36fun7", small: "_1j36fun8", caption: "_1j36fun9", body: "_1j36funa", lead: "_1j36funb", headline21: "_1j36func", headline26: "_1j36fund", headline31: "_1j36fune", headline37: "_1j36funf", headline48: "_1j36fung", display: "_1j36funh", hero: "_1j36funi", uber: "_1j36funj", colossus: "_1j36funk" }, type: { regular: "_1j36funl", bold: "_1j36funm" } }, defaultVariants: { type: "regular" }, compoundVariants: [] });
var button_css_ts_vanilla = "";
var buttonStyle = createRuntimeFn({ defaultClassName: "_164adwm0", variantClassNames: { size: { medium: "_164adwm1", small: "_164adwm2", large: "_164adwm3" }, variant: { regular: "_164adwm4", semibold: "_164adwm5" } }, defaultVariants: {}, compoundVariants: [] });
var buttonLink_css_ts_vanilla = "";
var buttonLinkStyle = createRuntimeFn({ defaultClassName: "_1wdteob0", variantClassNames: {}, defaultVariants: {}, compoundVariants: [] });
var buttonSocial_css_ts_vanilla = "";
var buttonSocialIconStyle = createRuntimeFn({ defaultClassName: "_1ddohgb7", variantClassNames: { size: { small: "_1ddohgb8", medium: "_1ddohgb9", large: "_1ddohgba" } }, defaultVariants: {}, compoundVariants: [] });
var buttonSocialStyle = createRuntimeFn({ defaultClassName: "_1ddohgb0", variantClassNames: { size: { small: "_1ddohgb1", medium: "_1ddohgb2", large: "_1ddohgb3" }, variant: { regular: "_1ddohgb4", semibold: "_1ddohgb5" } }, defaultVariants: {}, compoundVariants: [] });
var buttonSocialTitleStyle = "_1ddohgb6";
var colors_css_ts_vanilla = "";
function _defineProperty(obj, key, value) {
  if (key in obj) {
    Object.defineProperty(obj, key, {
      value,
      enumerable: true,
      configurable: true,
      writable: true
    });
  } else {
    obj[key] = value;
  }
  return obj;
}
function ownKeys(object, enumerableOnly) {
  var keys = Object.keys(object);
  if (Object.getOwnPropertySymbols) {
    var symbols = Object.getOwnPropertySymbols(object);
    if (enumerableOnly) {
      symbols = symbols.filter(function(sym) {
        return Object.getOwnPropertyDescriptor(object, sym).enumerable;
      });
    }
    keys.push.apply(keys, symbols);
  }
  return keys;
}
function _objectSpread2(target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i] != null ? arguments[i] : {};
    if (i % 2) {
      ownKeys(Object(source), true).forEach(function(key) {
        _defineProperty(target, key, source[key]);
      });
    } else if (Object.getOwnPropertyDescriptors) {
      Object.defineProperties(target, Object.getOwnPropertyDescriptors(source));
    } else {
      ownKeys(Object(source)).forEach(function(key) {
        Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key));
      });
    }
  }
  return target;
}
var createSprinkles$1 = (composeStyles2) => function() {
  for (var _len = arguments.length, args = new Array(_len), _key = 0; _key < _len; _key++) {
    args[_key] = arguments[_key];
  }
  var sprinklesStyles = Object.assign({}, ...args.map((a) => a.styles));
  var sprinklesKeys = Object.keys(sprinklesStyles);
  var shorthandNames = sprinklesKeys.filter((property) => "mappings" in sprinklesStyles[property]);
  var sprinklesFn = (props) => {
    var classNames = [];
    var shorthands = {};
    var nonShorthands = _objectSpread2({}, props);
    var hasShorthands = false;
    for (var shorthand of shorthandNames) {
      var value = props[shorthand];
      if (value != null) {
        var sprinkle = sprinklesStyles[shorthand];
        hasShorthands = true;
        for (var propMapping of sprinkle.mappings) {
          shorthands[propMapping] = value;
          if (nonShorthands[propMapping] == null) {
            delete nonShorthands[propMapping];
          }
        }
      }
    }
    var finalProps = hasShorthands ? _objectSpread2(_objectSpread2({}, shorthands), nonShorthands) : props;
    for (var prop in finalProps) {
      var propValue = finalProps[prop];
      var _sprinkle = sprinklesStyles[prop];
      try {
        if (_sprinkle.mappings) {
          continue;
        }
        if (typeof propValue === "string" || typeof propValue === "number") {
          if (false)
            ;
          classNames.push(_sprinkle.values[propValue].defaultClass);
        } else if (Array.isArray(propValue)) {
          for (var responsiveIndex = 0; responsiveIndex < propValue.length; responsiveIndex++) {
            var responsiveValue = propValue[responsiveIndex];
            if (responsiveValue != null) {
              var conditionName = _sprinkle.responsiveArray[responsiveIndex];
              if (false)
                ;
              classNames.push(_sprinkle.values[responsiveValue].conditions[conditionName]);
            }
          }
        } else {
          for (var _conditionName in propValue) {
            var _value = propValue[_conditionName];
            if (_value != null) {
              if (false)
                ;
              classNames.push(_sprinkle.values[_value].conditions[_conditionName]);
            }
          }
        }
      } catch (e) {
        throw e;
      }
    }
    return composeStyles2(classNames.join(" "));
  };
  return Object.assign(sprinklesFn, {
    properties: new Set(sprinklesKeys)
  });
};
var composeStyles = (classList) => classList;
var createSprinkles = function createSprinkles2() {
  return createSprinkles$1(composeStyles)(...arguments);
};
var colorProperties = { conditions: void 0, styles: { color: { values: { accentDisabled: { defaultClass: "bf9b9v0" }, accentDefault: { defaultClass: "bf9b9v1" }, accentMuted: { defaultClass: "bf9b9v2" }, accentSubtle: { defaultClass: "bf9b9v3" }, accentEmphasis: { defaultClass: "bf9b9v4" }, foregroundDefault: { defaultClass: "bf9b9v5" }, foregroundMuted: { defaultClass: "bf9b9v6" }, foregroundSubtle: { defaultClass: "bf9b9v7" }, foregroundDisabled: { defaultClass: "bf9b9v8" }, foregroundOnDark: { defaultClass: "bf9b9v9" }, foregroundOnAccent: { defaultClass: "bf9b9va" }, backgroundSurface: { defaultClass: "bf9b9vb" }, backgroundCanvas: { defaultClass: "bf9b9vc" }, errorDefault: { defaultClass: "bf9b9vd" }, errorSubtle: { defaultClass: "bf9b9ve" }, errorMuted: { defaultClass: "bf9b9vf" }, errorEmphasis: { defaultClass: "bf9b9vg" }, successEmphasis: { defaultClass: "bf9b9vh" }, borderDefault: { defaultClass: "bf9b9vi" }, textDefault: { defaultClass: "bf9b9vj" }, textDisabled: { defaultClass: "bf9b9vk" }, inputBackground: { defaultClass: "bf9b9vl" }, inputDisabled: { defaultClass: "bf9b9vm" }, inputPlaceholder: { defaultClass: "bf9b9vn" }, inputText: { defaultClass: "bf9b9vo" } } } } };
var colorSprinkle = createSprinkles({ conditions: void 0, styles: { color: { values: { accentDisabled: { defaultClass: "bf9b9v0" }, accentDefault: { defaultClass: "bf9b9v1" }, accentMuted: { defaultClass: "bf9b9v2" }, accentSubtle: { defaultClass: "bf9b9v3" }, accentEmphasis: { defaultClass: "bf9b9v4" }, foregroundDefault: { defaultClass: "bf9b9v5" }, foregroundMuted: { defaultClass: "bf9b9v6" }, foregroundSubtle: { defaultClass: "bf9b9v7" }, foregroundDisabled: { defaultClass: "bf9b9v8" }, foregroundOnDark: { defaultClass: "bf9b9v9" }, foregroundOnAccent: { defaultClass: "bf9b9va" }, backgroundSurface: { defaultClass: "bf9b9vb" }, backgroundCanvas: { defaultClass: "bf9b9vc" }, errorDefault: { defaultClass: "bf9b9vd" }, errorSubtle: { defaultClass: "bf9b9ve" }, errorMuted: { defaultClass: "bf9b9vf" }, errorEmphasis: { defaultClass: "bf9b9vg" }, successEmphasis: { defaultClass: "bf9b9vh" }, borderDefault: { defaultClass: "bf9b9vi" }, textDefault: { defaultClass: "bf9b9vj" }, textDisabled: { defaultClass: "bf9b9vk" }, inputBackground: { defaultClass: "bf9b9vl" }, inputDisabled: { defaultClass: "bf9b9vm" }, inputPlaceholder: { defaultClass: "bf9b9vn" }, inputText: { defaultClass: "bf9b9vo" } } } } });
var oryTheme = { fontFamily: "var(--ory-theme-font-family)", fontStyle: "var(--ory-theme-font-style)", accent: { def: "var(--ory-theme-accent-def)", muted: "var(--ory-theme-accent-muted)", emphasis: "var(--ory-theme-accent-emphasis)", disabled: "var(--ory-theme-accent-disabled)", subtle: "var(--ory-theme-accent-subtle)" }, foreground: { def: "var(--ory-theme-foreground-def)", muted: "var(--ory-theme-foreground-muted)", subtle: "var(--ory-theme-foreground-subtle)", disabled: "var(--ory-theme-foreground-disabled)", onDark: "var(--ory-theme-foreground-on-dark)", onAccent: "var(--ory-theme-foreground-on-accent)", onDisabled: "var(--ory-theme-foreground-on-disabled)" }, background: { surface: "var(--ory-theme-background-surface)", canvas: "var(--ory-theme-background-canvas)" }, error: { def: "var(--ory-theme-error-def)", subtle: "var(--ory-theme-error-subtle)", muted: "var(--ory-theme-error-muted)", emphasis: "var(--ory-theme-error-emphasis)" }, success: { emphasis: "var(--ory-theme-success-emphasis)" }, border: { def: "var(--ory-theme-border-def)" }, text: { def: "var(--ory-theme-text-def)", disabled: "var(--ory-theme-text-disabled)" }, input: { background: "var(--ory-theme-input-background)", disabled: "var(--ory-theme-input-disabled)", placeholder: "var(--ory-theme-input-placeholder)", text: "var(--ory-theme-input-text)" } };
var grid_css_ts_vanilla = "";
var gridStyle = createRuntimeFn({ defaultClassName: "juizgb0", variantClassNames: { direction: { row: "juizgb1", column: "juizgb2" }, gap: { "4": "juizgb3", "8": "juizgb4", "16": "juizgb5", "32": "juizgb6", "64": "juizgb7" } }, defaultVariants: {}, compoundVariants: [] });
var checkbox_css_ts_vanilla = "";
var checkboxInputStyle = "_1m7dk2l1";
var checkboxStyle = "_1m7dk2l0";
var inputField_css_ts_vanilla = "";
var inputFieldStyle = "a0s6411";
var inputFieldTitleStyle = "a0s6410";
export { buttonLinkStyle, buttonSocialIconStyle, buttonSocialStyle, buttonSocialTitleStyle, buttonStyle, cardStyle, cardTitleStyle, checkboxInputStyle, checkboxStyle, colorProperties, colorSprinkle, defaultBreakpoints, defaultDarkTheme, defaultFont, defaultLightTheme, dividerStyle, gridStyle, inputFieldStyle, inputFieldTitleStyle, inputTypographyStyle, messageStyle, oryTheme, typographyStyle };
//# sourceMappingURL=index.es.js.map
