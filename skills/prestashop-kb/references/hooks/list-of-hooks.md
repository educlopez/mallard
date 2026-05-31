---
source_url: https://devdocs.prestashop-project.org/9/modules/concepts/hooks/list-of-hooks/
ps_version: [8, 9]
ingested: 2026-04-18
section: Hooks > List of all available hooks
topic: hooks
tags: [prestashop, hooks, list, reference, display, action]
---

[Improve this page](https://github.com/PrestaShop/docs/edit/9.x/modules/concepts/hooks/list-of-hooks/_index.md "Improve this page") [How](/9/contribute/documentation/how)

**Search tip:** Some hooks are generated dynamically, so their names are documented in a generic way.

For example, `actionAdminCustomersFormModifier` is documented as `action<AdminControllerClassName>FormModifier`. A regex based search has been implemented, and generic hooks should be matched. However, when you see a controller name or action in the hook name and you can’t find it, try searching for a part of the hook name, like `FormModifier`.

Looking for available hooks in the front office? Check out our [Hook mapping project](/9/themes/hummingbird/hooks/), which provides a visual representation of hook availability in the Hummingbird theme.

## Search for a hook

No hooks found

| Locations | Example | Name | Aliases | Title | Type |
| --- | --- | --- | --- | --- | --- |
| BO |  | [action<ClassName><Action>After](/9/modules/concepts/hooks/list-of-hooks/actionclassnameactionafter/) |  |  | action |
| BO |  | [action<ClassName><Action>Before](/9/modules/concepts/hooks/list-of-hooks/actionclassnameactionbefore/) |  |  | action |
| BO |  | [action<Controller>FormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontrollerformmodifier/) |  |  | action |
| FO |  | [action<Controller>InitAfter](/9/modules/concepts/hooks/list-of-hooks/actioncontrollerinitafter/) |  |  | action |
| FO |  | [action<Controller>InitBefore](/9/modules/concepts/hooks/list-of-hooks/actioncontrollerinitbefore/) |  |  | action |
| BO |  | [action<Controller>ListingFieldsModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontrollerlistingfieldsmodifier/) |  |  | action |
| BO |  | [action<Controller>ListingResultsModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontrollerlistingresultsmodifier/) |  |  | action |
| BO |  | [action<Controller>OptionsModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontrolleroptionsmodifier/) |  |  | action |
| FO |  | [action<Controller>SetMedia](/9/modules/concepts/hooks/list-of-hooks/actioncontrollersetmedia/) |  |  | action |
| FO |  | [action<Controller>SetVariables](/9/modules/concepts/hooks/list-of-hooks/actioncontrollersetvariables/) |  |  | action |
| FO |  | [action<Controller>SetVariablesBefore](/9/modules/concepts/hooks/list-of-hooks/actioncontrollersetvariablesbefore/) |  |  | action |
| FO  BO |  | [action<DefinitionId>GridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actiondefinitionidgriddatamodifier/) |  |  | action |
| FO  BO | ✅ | [action<DefinitionId>GridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actiondefinitionidgriddefinitionmodifier/) |  |  | action |
| FO  BO |  | [action<DefinitionId>GridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actiondefinitionidgridfilterformmodifier/) |  |  | action |
| FO |  | [action<DefinitionId>GridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actiondefinitionidgridpresentermodifier/) |  |  | action |
| FO  BO | ✅ | [action<DefinitionId>GridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actiondefinitionidgridquerybuildermodifier/) |  |  | action |
| FO  BO | ✅ | [action<FormName>FormBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionformnameformbuildermodifier/) |  |  | action |
| FO  BO |  | [action<FormName>FormDataProviderData](/9/modules/concepts/hooks/list-of-hooks/actionformnameformdataproviderdata/) |  |  | action |
| FO |  | [action<FormName>FormDataProviderDefaultData](/9/modules/concepts/hooks/list-of-hooks/actionformnameformdataproviderdefaultdata/) |  |  | action |
| FO |  | [action<LegacyControllerName>ListingFieldsModifier](/9/modules/concepts/hooks/list-of-hooks/actionlegacycontrollernamelistingfieldsmodifier/) |  |  | action |
| FO |  | [action<LegacyControllerName>ListingResultsModifier](/9/modules/concepts/hooks/list-of-hooks/actionlegacycontrollernamelistingresultsmodifier/) |  |  | action |
| BO |  | [actionAddressGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionaddressgriddatamodifier/) |  | Modify address grid data | action |
| BO |  | [actionAddressGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionaddressgriddefinitionmodifier/) |  | Modify address grid definition | action |
| BO |  | [actionAddressGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionaddressgridfilterformmodifier/) |  | Modify address grid filters | action |
| BO |  | [actionAddressGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionaddressgridpresentermodifier/) |  | Modify address grid template data | action |
| BO |  | [actionAddressGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionaddressgridquerybuildermodifier/) |  | Modify address grid query builder | action |
| BO |  | [actionAdmin<Action>After](/9/modules/concepts/hooks/list-of-hooks/actionadminactionafter/) |  |  | action |
| BO |  | [actionAdmin<Action>Before](/9/modules/concepts/hooks/list-of-hooks/actionadminactionbefore/) |  |  | action |
| BO |  | [actionAdminAdministrationControllerPostProcess<HookName>Before](/9/modules/concepts/hooks/list-of-hooks/actionadminadministrationcontrollerpostprocesshooknamebefore/) |  |  | action |
| BO |  | [actionAdminAdministrationControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminadministrationcontrollerpostprocessbefore/) |  | On post-process in Admin Configure Advanced Parameters Administration Controller | action |
| BO |  | [actionAdminAdminShopParametersMetaControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminadminshopparametersmetacontrollerpostprocessbefore/) |  | On post-process in Admin Configure Shop Parameters Meta Controller | action |
| BO |  | [actionAdminAdminWebserviceControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminadminwebservicecontrollerpostprocessbefore/) |  |  | action |
| BO |  | [actionAdminAdvancedParametersPerformanceControllerPostProcess<HookName>Before](/9/modules/concepts/hooks/list-of-hooks/actionadminadvancedparametersperformancecontrollerpostprocesshooknamebefore/) |  |  | action |
| BO |  | [actionAdminAdvancedParametersPerformanceControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminadvancedparametersperformancecontrollerpostprocessbefore/) |  | On post-process in Admin Configure Advanced Parameters Performance Controller | action |
| BO |  | [actionAdminBreadcrumbModifier](/9/modules/concepts/hooks/list-of-hooks/actionadminbreadcrumbmodifier/) |  | Modify back office breadcrumb | action |
| BO |  | [actionAdminControllerInitAfter](/9/modules/concepts/hooks/list-of-hooks/actionadmincontrollerinitafter/) |  | Perform actions after admin controller initialization | action |
| BO |  | [actionAdminControllerInitBefore](/9/modules/concepts/hooks/list-of-hooks/actionadmincontrollerinitbefore/) |  | Perform actions before admin controller initialization | action |
| BO | ✅ | [actionAdminControllerSetMedia](/9/modules/concepts/hooks/list-of-hooks/actionadmincontrollersetmedia/) |  |  | action |
| BO |  | [actionAdminDuplicateAfter](/9/modules/concepts/hooks/list-of-hooks/actionadminduplicateafter/) |  |  | action |
| BO |  | [actionAdminDuplicateBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminduplicatebefore/) |  |  | action |
| BO |  | [actionAdminDuplicateDiscountAfter](/9/modules/concepts/hooks/list-of-hooks/actionadminduplicatediscountafter/) |  |  | action |
| BO |  | [actionAdminDuplicateDiscountBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminduplicatediscountbefore/) |  |  | action |
| BO |  | [actionAdminInternationalGeolocationControllerPostProcess<HookName>Before](/9/modules/concepts/hooks/list-of-hooks/actionadmininternationalgeolocationcontrollerpostprocesshooknamebefore/) |  |  | action |
| BO |  | [actionAdminInternationalGeolocationControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadmininternationalgeolocationcontrollerpostprocessbefore/) |  | On post-process in Admin Improve International Geolocation Controller | action |
| BO |  | [actionAdminInternationalLocalizationControllerPostProcess<HookName>Before](/9/modules/concepts/hooks/list-of-hooks/actionadmininternationallocalizationcontrollerpostprocesshooknamebefore/) |  |  | action |
| BO |  | [actionAdminInternationalLocalizationControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadmininternationallocalizationcontrollerpostprocessbefore/) |  | On post-process in Admin Improve International Localization Controller | action |
| BO |  | [actionAdminLoginControllerSetMedia](/9/modules/concepts/hooks/list-of-hooks/actionadminlogincontrollersetmedia/) |  | Set media on admin login page header | action |
| BO |  | [actionAdminLogsControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminlogscontrollerpostprocessbefore/) |  |  | action |
| BO |  | [actionAdminMaintenanceControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminmaintenancecontrollerpostprocessbefore/) |  |  | action |
| BO |  | [actionAdminMenuTabsModifier](/9/modules/concepts/hooks/list-of-hooks/actionadminmenutabsmodifier/) |  | Modify back office menu | action |
| BO |  | [actionAdminMetaAfterWriteRobotsFile](/9/modules/concepts/hooks/list-of-hooks/actionadminmetaafterwriterobotsfile/) |  |  | action |
| BO |  | [actionAdminMetaBeforeWriteRobotsFile](/9/modules/concepts/hooks/list-of-hooks/actionadminmetabeforewriterobotsfile/) |  |  | action |
| BO |  | [actionAdminMetaSave](/9/modules/concepts/hooks/list-of-hooks/actionadminmetasave/) | afterSaveAdminMeta | After saving the configuration in AdminMeta | action |
| BO |  | [actionAdminOrdersTrackingNumberUpdate](/9/modules/concepts/hooks/list-of-hooks/actionadminorderstrackingnumberupdate/) |  | After setting the tracking number for the order | action |
| BO |  | [actionAdminPreferencesControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminpreferencescontrollerpostprocessbefore/) |  |  | action |
| BO |  | [actionAdminProductsListingFieldsModifier](/9/modules/concepts/hooks/list-of-hooks/actionadminproductslistingfieldsmodifier/) |  |  | action |
| BO |  | [actionAdminProductsListingResultsModifier](/9/modules/concepts/hooks/list-of-hooks/actionadminproductslistingresultsmodifier/) |  |  | action |
| BO |  | [actionAdminSecurityControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminsecuritycontrollerpostprocessbefore/) |  | On post-process in Admin Security Controller | action |
| BO |  | [actionAdminShippingPreferencesControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminshippingpreferencescontrollerpostprocessbefore/) |  | On post-process in Admin Improve Shipping Preferences Controller | action |
| BO |  | [actionAdminShippingPreferencesControllerPostProcessCarrierOptionsBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminshippingpreferencescontrollerpostprocesscarrieroptionsbefore/) |  | On post-process in Admin Improve Shipping Preferences Controller | action |
| BO |  | [actionAdminShippingPreferencesControllerPostProcessHandlingBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminshippingpreferencescontrollerpostprocesshandlingbefore/) |  | On post-process in Admin Improve Shipping Preferences Controller | action |
| BO |  | [actionAdminShopParametersMetaControllerPostProcess<HookName>Before](/9/modules/concepts/hooks/list-of-hooks/actionadminshopparametersmetacontrollerpostprocesshooknamebefore/) |  |  | action |
| BO |  | [actionAdminShopParametersOrderPreferencesControllerPostProcess<HookName>Before](/9/modules/concepts/hooks/list-of-hooks/actionadminshopparametersorderpreferencescontrollerpostprocesshooknamebefore/) |  |  | action |
| BO |  | [actionAdminShopParametersOrderPreferencesControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminshopparametersorderpreferencescontrollerpostprocessbefore/) |  | On post-process in Admin Configure Shop Parameters Order Preferences Controller | action |
| BO |  | [actionAdminShopParametersProductPreferencesControllerPostProcess<HookName>Before](/9/modules/concepts/hooks/list-of-hooks/actionadminshopparametersproductpreferencescontrollerpostprocesshooknamebefore/) |  |  | action |
| BO |  | [actionAdminShopParametersProductPreferencesControllerPostProcessBefore](/9/modules/concepts/hooks/list-of-hooks/actionadminshopparametersproductpreferencescontrollerpostprocessbefore/) |  |  | action |
| BO |  | [actionAdminThemesControllerUpdate\_optionsAfter](/9/modules/concepts/hooks/list-of-hooks/actionadminthemescontrollerupdate_optionsafter/) |  |  | action |
| BO |  | [actionAdminThemesControllerUpdateoptionsAfter](/9/modules/concepts/hooks/list-of-hooks/actionadminthemescontrollerupdateoptionsafter/) |  |  | action |
| BO | ✅ | [actionAfterCreate<FormName>FormHandler](/9/modules/concepts/hooks/list-of-hooks/actionaftercreateformnameformhandler/) |  |  | action |
| FO  BO |  | [actionAfterLoadRoutes](/9/modules/concepts/hooks/list-of-hooks/actionafterloadroutes/) |  |  | action |
| BO | ✅ | [actionAfterUpdate<FormName>FormHandler](/9/modules/concepts/hooks/list-of-hooks/actionafterupdateformnameformhandler/) |  |  | action |
| FO |  | [actionAjaxDie<Controller><Method>Before](/9/modules/concepts/hooks/list-of-hooks/actionajaxdiecontrollermethodbefore/) |  |  | action |
| FO |  | [actionApplyCartRule](/9/modules/concepts/hooks/list-of-hooks/actionapplycartrule/) |  |  | action |
| BO |  | [actionAttachmentGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionattachmentgriddatamodifier/) |  | Modify attachment grid data | action |
| BO |  | [actionAttachmentGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionattachmentgriddefinitionmodifier/) |  | Modify attachment grid definition | action |
| BO |  | [actionAttachmentGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionattachmentgridfilterformmodifier/) |  | Modify attachment grid filters | action |
| BO |  | [actionAttachmentGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionattachmentgridpresentermodifier/) |  | Modify attachment grid template data | action |
| BO |  | [actionAttachmentGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionattachmentgridquerybuildermodifier/) |  | Modify attachment grid query builder | action |
| FO |  | [actionAttributeCombinationDelete](/9/modules/concepts/hooks/list-of-hooks/actionattributecombinationdelete/) |  |  | action |
| FO |  | [actionAttributeCombinationSave](/9/modules/concepts/hooks/list-of-hooks/actionattributecombinationsave/) |  |  | action |
| FO |  | [actionAttributeDelete](/9/modules/concepts/hooks/list-of-hooks/actionattributedelete/) | afterDeleteAttribute | Deleting an attributes features value | action |
| FO |  | [actionAttributeGroupDelete](/9/modules/concepts/hooks/list-of-hooks/actionattributegroupdelete/) | afterDeleteAttributeGroup | Deleting attribute group | action |
| BO |  | [actionAttributeGroupGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionattributegroupgriddatamodifier/) |  | Modify attribute group grid data | action |
| BO |  | [actionAttributeGroupGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionattributegroupgriddefinitionmodifier/) |  | Modify attribute group grid definition | action |
| BO |  | [actionAttributeGroupGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionattributegroupgridfilterformmodifier/) |  | Modify attribute group grid filters | action |
| BO |  | [actionAttributeGroupGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionattributegroupgridpresentermodifier/) |  | Modify attribute group grid template data | action |
| BO |  | [actionAttributeGroupGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionattributegroupgridquerybuildermodifier/) |  | Modify attribute group grid query builder | action |
| FO |  | [actionAttributeGroupSave](/9/modules/concepts/hooks/list-of-hooks/actionattributegroupsave/) | afterSaveAttributeGroup | Saving an attribute group | action |
| FO |  | [actionAttributeSave](/9/modules/concepts/hooks/list-of-hooks/actionattributesave/) | afterSaveAttribute | Saving an attributes features value | action |
| FO |  | [actionAuthentication](/9/modules/concepts/hooks/list-of-hooks/actionauthentication/) | authentication | Successful customer authentication | action |
| FO |  | [actionAuthenticationBefore](/9/modules/concepts/hooks/list-of-hooks/actionauthenticationbefore/) | actionBeforeAuthentication |  | action |
| BO |  | [actionBackupGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionbackupgriddatamodifier/) |  | Modify backup grid data | action |
| BO |  | [actionBackupGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionbackupgriddefinitionmodifier/) |  | Modifying DB Backup grid definition | action |
| BO |  | [actionBackupGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionbackupgridfilterformmodifier/) |  | Modify filters form for DB Backup grid | action |
| BO |  | [actionBackupGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionbackupgridpresentermodifier/) |  | Modify DB Backup grid view data | action |
| BO |  | [actionBackupGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionbackupgridquerybuildermodifier/) |  | Modify backup grid query builder | action |
| FO |  | [actionBeforeAjaxDie<Controller><Method>](/9/modules/concepts/hooks/list-of-hooks/actionbeforeajaxdiecontrollermethod/) |  |  | action |
| FO |  | [actionBeforeCreate<FormName>FormHandler](/9/modules/concepts/hooks/list-of-hooks/actionbeforecreateformnameformhandler/) |  |  | action |
| FO |  | [actionBeforeUpdate<FormName>FormHandler](/9/modules/concepts/hooks/list-of-hooks/actionbeforeupdateformnameformhandler/) |  |  | action |
| FO |  | [actionBuild<Controller>FrontEndObject](/9/modules/concepts/hooks/list-of-hooks/actionbuildcontrollerfrontendobject/) |  |  | action |
| FO |  | [actionBuildFrontEndObject](/9/modules/concepts/hooks/list-of-hooks/actionbuildfrontendobject/) |  | Manage elements added to the "prestashop" javascript object | action |
| FO | ✅ | [actionBuildMailLayoutVariables](/9/modules/concepts/hooks/list-of-hooks/actionbuildmaillayoutvariables/) |  | Build the variables used in email layout rendering | action |
| BO |  | [actionCarrierGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncarriergriddatamodifier/) |  | Modify carrier grid data | action |
| BO |  | [actionCarrierGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncarriergriddefinitionmodifier/) |  | Modify carrier grid definition | action |
| BO |  | [actionCarrierGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncarriergridfilterformmodifier/) |  | Modify carrier grid filters | action |
| BO |  | [actionCarrierGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncarriergridpresentermodifier/) |  | Modify carrier grid template data | action |
| BO |  | [actionCarrierGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncarriergridquerybuildermodifier/) |  | Modify carrier grid query builder | action |
| FO |  | [actionCarrierProcess](/9/modules/concepts/hooks/list-of-hooks/actioncarrierprocess/) | processCarrier | Carrier process | action |
| BO |  | [actionCarrierUpdate](/9/modules/concepts/hooks/list-of-hooks/actioncarrierupdate/) | updateCarrier | Carrier Update | action |
| FO |  | [actionCartGetPackageShippingCost](/9/modules/concepts/hooks/list-of-hooks/actioncartgetpackageshippingcost/) |  | After getting package shipping cost value | action |
| BO |  | [actionCartRuleGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncartrulegriddatamodifier/) |  | Modify cart rule grid data | action |
| BO |  | [actionCartRuleGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncartrulegriddefinitionmodifier/) |  | Modify cart rule grid definition | action |
| BO |  | [actionCartRuleGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncartrulegridfilterformmodifier/) |  | Modify cart rule grid filters | action |
| BO |  | [actionCartRuleGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncartrulegridpresentermodifier/) |  | Modify cart rule grid template data | action |
| BO |  | [actionCartRuleGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncartrulegridquerybuildermodifier/) |  | Modify cart rule grid query builder | action |
| FO |  | [actionCartSave](/9/modules/concepts/hooks/list-of-hooks/actioncartsave/) | cart | Cart creation and update | action |
| FO |  | [actionCartSummary](/9/modules/concepts/hooks/list-of-hooks/actioncartsummary/) |  |  | action |
| FO |  | [actionCartUpdateQuantityBefore](/9/modules/concepts/hooks/list-of-hooks/actioncartupdatequantitybefore/) | actionBeforeCartUpdateQty |  | action |
| BO |  | [actionCatalogPriceRuleGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncatalogpricerulegriddatamodifier/) |  | Modify catalog price rule grid data | action |
| BO |  | [actionCatalogPriceRuleGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncatalogpricerulegriddefinitionmodifier/) |  | Modify catalog price rule grid definition | action |
| BO |  | [actionCatalogPriceRuleGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncatalogpricerulegridfilterformmodifier/) |  | Modify catalog price rule grid filters | action |
| BO |  | [actionCatalogPriceRuleGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncatalogpricerulegridpresentermodifier/) |  | Modify catalog price rule grid template data | action |
| BO |  | [actionCatalogPriceRuleGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncatalogpricerulegridquerybuildermodifier/) |  | Modify catalog price rule grid query builder | action |
| FO |  | [actionCategoryAdd](/9/modules/concepts/hooks/list-of-hooks/actioncategoryadd/) | categoryAddition | Category creation | action |
| FO |  | [actionCategoryDelete](/9/modules/concepts/hooks/list-of-hooks/actioncategorydelete/) | categoryDeletion | Category deletion | action |
| BO |  | [actionCategoryGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncategorygriddatamodifier/) |  | Modify category grid data | action |
| BO |  | [actionCategoryGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncategorygriddefinitionmodifier/) |  | Modify category grid definition | action |
| BO |  | [actionCategoryGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncategorygridfilterformmodifier/) |  | Modify category grid filters | action |
| BO |  | [actionCategoryGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncategorygridpresentermodifier/) |  | Modify category grid template data | action |
| BO |  | [actionCategoryGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncategorygridquerybuildermodifier/) |  | Modify category grid query builder | action |
| BO |  | [actionCategoryUpdate](/9/modules/concepts/hooks/list-of-hooks/actioncategoryupdate/) | categoryUpdate | Category modification | action |
| FO |  | [actionCheckAttributeQuantity](/9/modules/concepts/hooks/list-of-hooks/actioncheckattributequantity/) |  | Check product attribute quantity availability | action |
| FO |  | [actionCheckoutRender](/9/modules/concepts/hooks/list-of-hooks/actioncheckoutrender/) |  | Modify checkout process | action |
| FO |  | [actionCheckoutStepRenderTemplate](/9/modules/concepts/hooks/list-of-hooks/actioncheckoutsteprendertemplate/) |  | Modify the parameters of the checkout step template rendering | action |
| FO |  | [actionClearCache](/9/modules/concepts/hooks/list-of-hooks/actionclearcache/) |  | Clear smarty cache | action |
| FO |  | [actionClearCompileCache](/9/modules/concepts/hooks/list-of-hooks/actionclearcompilecache/) |  | Clear smarty compile cache | action |
| FO |  | [actionClearSf2Cache](/9/modules/concepts/hooks/list-of-hooks/actionclearsf2cache/) |  | Clear Sf2 cache | action |
| BO |  | [actionCmsPageCategoryGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagecategorygriddatamodifier/) |  | Modify cms page category grid data | action |
| BO |  | [actionCmsPageCategoryGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagecategorygriddefinitionmodifier/) |  | Modify cms page category grid definition | action |
| BO |  | [actionCmsPageCategoryGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagecategorygridfilterformmodifier/) |  | Modify cms page category grid filters | action |
| BO |  | [actionCmsPageCategoryGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagecategorygridpresentermodifier/) |  | Modify cms page category grid template data | action |
| BO |  | [actionCmsPageCategoryGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagecategorygridquerybuildermodifier/) |  | Modify cms page category grid query builder | action |
| BO |  | [actionCmsPageGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagegriddatamodifier/) |  | Modify cms page grid data | action |
| BO |  | [actionCmsPageGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagegriddefinitionmodifier/) |  | Modify cms page grid definition | action |
| BO |  | [actionCmsPageGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagegridfilterformmodifier/) |  | Modify cms page grid filters | action |
| BO |  | [actionCmsPageGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagegridpresentermodifier/) |  | Modify cms page grid template data | action |
| BO |  | [actionCmsPageGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncmspagegridquerybuildermodifier/) |  | Modify cms page grid query builder | action |
| FO |  | [actionConfigurationUpdateValueBefore](/9/modules/concepts/hooks/list-of-hooks/actionconfigurationupdatevaluebefore/) |  |  | action |
| BO |  | [actionContactGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontactgriddatamodifier/) |  | Modify contact grid data | action |
| BO |  | [actionContactGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontactgriddefinitionmodifier/) |  | Modify contact grid definition | action |
| BO |  | [actionContactGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontactgridfilterformmodifier/) |  | Modify contact grid filters | action |
| BO |  | [actionContactGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontactgridpresentermodifier/) |  | Modify contact grid template data | action |
| BO |  | [actionContactGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncontactgridquerybuildermodifier/) |  | Modify contact grid query builder | action |
| FO |  | [actionControllerInitAfter](/9/modules/concepts/hooks/list-of-hooks/actioncontrollerinitafter/) |  | Perform actions after controller initialization | action |
| FO |  | [actionControllerInitBefore](/9/modules/concepts/hooks/list-of-hooks/actioncontrollerinitbefore/) |  | Perform actions before controller initialization | action |
| BO |  | [actionCountryGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncountrygriddatamodifier/) |  | Modify country grid data | action |
| BO |  | [actionCountryGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncountrygriddefinitionmodifier/) |  | Modify country grid definition | action |
| BO |  | [actionCountryGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncountrygridfilterformmodifier/) |  | Modify country grid filters | action |
| BO |  | [actionCountryGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncountrygridpresentermodifier/) |  | Modify country grid template data | action |
| BO |  | [actionCountryGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncountrygridquerybuildermodifier/) |  | Modify country grid query builder | action |
| BO |  | [actionCreditSlipGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncreditslipgriddatamodifier/) |  | Modify credit slip grid data | action |
| BO |  | [actionCreditSlipGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncreditslipgriddefinitionmodifier/) |  | Modify credit slip grid definition | action |
| BO |  | [actionCreditSlipGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncreditslipgridfilterformmodifier/) |  | Modify credit slip grid filters | action |
| BO |  | [actionCreditSlipGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncreditslipgridpresentermodifier/) |  | Modify credit slip grid template data | action |
| BO |  | [actionCreditSlipGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncreditslipgridquerybuildermodifier/) |  | Modify credit slip grid query builder | action |
| BO |  | [actionCurrencyGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncurrencygriddatamodifier/) |  | Modify currency grid data | action |
| BO |  | [actionCurrencyGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncurrencygriddefinitionmodifier/) |  | Modify currency grid definition | action |
| BO |  | [actionCurrencyGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncurrencygridfilterformmodifier/) |  | Modify currency grid filters | action |
| BO |  | [actionCurrencyGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncurrencygridpresentermodifier/) |  | Modify currency grid template data | action |
| BO |  | [actionCurrencyGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncurrencygridquerybuildermodifier/) |  | Modify currency grid query builder | action |
| FO |  | [actionCustomerAccountAdd](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraccountadd/) | createAccount | Successful customer account creation | action |
| FO |  | [actionCustomerAccountUpdate](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraccountupdate/) |  | Successful customer account update | action |
| FO |  | [actionCustomerAddGroups](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraddgroups/) |  |  | action |
| BO |  | [actionCustomerAddressGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraddressgriddatamodifier/) |  | Modify customer address grid data | action |
| BO |  | [actionCustomerAddressGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraddressgriddefinitionmodifier/) |  | Modify customer address grid definition | action |
| BO |  | [actionCustomerAddressGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraddressgridfilterformmodifier/) |  | Modify customer address grid filters | action |
| BO |  | [actionCustomerAddressGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraddressgridpresentermodifier/) |  | Modify customer address grid template data | action |
| BO |  | [actionCustomerAddressGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomeraddressgridquerybuildermodifier/) |  | Modify customer address grid query builder | action |
| FO |  | [actionCustomerBeforeUpdateGroup](/9/modules/concepts/hooks/list-of-hooks/actioncustomerbeforeupdategroup/) |  |  | action |
| BO |  | [actionCustomerBoughtProductGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerboughtproductgriddatamodifier/) |  | Modify customer bought product grid data | action |
| BO |  | [actionCustomerBoughtProductGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerboughtproductgriddefinitionmodifier/) |  | Modify customer bought product grid definition | action |
| BO |  | [actionCustomerBoughtProductGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerboughtproductgridfilterformmodifier/) |  | Modify customer bought product grid filters | action |
| BO |  | [actionCustomerBoughtProductGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerboughtproductgridpresentermodifier/) |  | Modify customer bought product grid template data | action |
| BO |  | [actionCustomerBoughtProductGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerboughtproductgridquerybuildermodifier/) |  | Modify customer bought product grid query builder | action |
| BO |  | [actionCustomerCartGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomercartgriddatamodifier/) |  | Modify customer cart grid data | action |
| BO |  | [actionCustomerCartGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomercartgriddefinitionmodifier/) |  | Modify customer cart grid definition | action |
| BO |  | [actionCustomerCartGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomercartgridfilterformmodifier/) |  | Modify customer cart grid filters | action |
| BO |  | [actionCustomerCartGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomercartgridpresentermodifier/) |  | Modify customer cart grid template data | action |
| BO |  | [actionCustomerCartGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomercartgridquerybuildermodifier/) |  | Modify customer cart grid query builder | action |
| BO |  | [actionCustomerDiscountGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerdiscountgriddatamodifier/) |  | Modify customer discount grid data | action |
| BO |  | [actionCustomerDiscountGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerdiscountgriddefinitionmodifier/) |  | Modify customer discount grid definition | action |
| BO |  | [actionCustomerDiscountGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerdiscountgridfilterformmodifier/) |  | Modify customer discount grid filters | action |
| BO |  | [actionCustomerDiscountGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerdiscountgridpresentermodifier/) |  | Modify customer discount grid template data | action |
| BO |  | [actionCustomerDiscountGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerdiscountgridquerybuildermodifier/) |  | Modify customer discount grid query builder | action |
| BO |  | [actionCustomerGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergriddatamodifier/) |  | Modify customer grid data | action |
| BO |  | [actionCustomerGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergriddefinitionmodifier/) |  | Modify customer grid definition | action |
| BO |  | [actionCustomerGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergridfilterformmodifier/) |  | Modify customer grid filters | action |
| BO |  | [actionCustomerGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergridpresentermodifier/) |  | Modify customer grid template data | action |
| BO |  | [actionCustomerGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergridquerybuildermodifier/) |  | Modify customer grid query builder | action |
| BO |  | [actionCustomerGroupsGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergroupsgriddatamodifier/) |  | Modify customer groups grid data | action |
| BO |  | [actionCustomerGroupsGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergroupsgriddefinitionmodifier/) |  | Modify customer groups grid definition | action |
| BO |  | [actionCustomerGroupsGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergroupsgridfilterformmodifier/) |  | Modify customer groups grid filters | action |
| BO |  | [actionCustomerGroupsGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergroupsgridpresentermodifier/) |  | Modify customer groups grid template data | action |
| BO |  | [actionCustomerGroupsGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomergroupsgridquerybuildermodifier/) |  | Modify customer groups grid query builder | action |
| FO |  | [actionCustomerLogoutAfter](/9/modules/concepts/hooks/list-of-hooks/actioncustomerlogoutafter/) |  | After customer logout | action |
| FO |  | [actionCustomerLogoutBefore](/9/modules/concepts/hooks/list-of-hooks/actioncustomerlogoutbefore/) |  | Before customer logout | action |
| BO |  | [actionCustomerOrderGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerordergriddatamodifier/) |  | Modify customer order grid data | action |
| BO |  | [actionCustomerOrderGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerordergriddefinitionmodifier/) |  | Modify customer order grid definition | action |
| BO |  | [actionCustomerOrderGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerordergridfilterformmodifier/) |  | Modify customer order grid filters | action |
| BO |  | [actionCustomerOrderGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerordergridpresentermodifier/) |  | Modify customer order grid template data | action |
| BO |  | [actionCustomerOrderGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerordergridquerybuildermodifier/) |  | Modify customer order grid query builder | action |
| BO |  | [actionCustomerThreadGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerthreadgriddatamodifier/) |  | Modify customer thread grid data | action |
| BO |  | [actionCustomerThreadGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerthreadgriddefinitionmodifier/) |  | Modify customer thread grid definition | action |
| BO |  | [actionCustomerThreadGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerthreadgridfilterformmodifier/) |  | Modify customer thread grid filters | action |
| BO |  | [actionCustomerThreadGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerthreadgridpresentermodifier/) |  | Modify customer thread grid template data | action |
| BO |  | [actionCustomerThreadGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerthreadgridquerybuildermodifier/) |  | Modify customer thread grid query builder | action |
| BO |  | [actionCustomerViewedProductGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerviewedproductgriddatamodifier/) |  | Modify customer viewed product grid data | action |
| BO |  | [actionCustomerViewedProductGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerviewedproductgriddefinitionmodifier/) |  | Modify customer viewed product grid definition | action |
| BO |  | [actionCustomerViewedProductGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerviewedproductgridfilterformmodifier/) |  | Modify customer viewed product grid filters | action |
| BO |  | [actionCustomerViewedProductGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerviewedproductgridpresentermodifier/) |  | Modify customer viewed product grid template data | action |
| BO |  | [actionCustomerViewedProductGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actioncustomerviewedproductgridquerybuildermodifier/) |  | Modify customer viewed product grid query builder | action |
| FO |  | [actionDeleteGDPRCustomer](/9/modules/concepts/hooks/list-of-hooks/actiondeletegdprcustomer/) |  |  | action |
| FO |  | [actionDeliveryPriceByPrice](/9/modules/concepts/hooks/list-of-hooks/actiondeliverypricebyprice/) |  |  | action |
| FO |  | [actionDeliveryPriceByWeight](/9/modules/concepts/hooks/list-of-hooks/actiondeliverypricebyweight/) |  |  | action |
| BO |  | [actionDisabledProductGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actiondisabledproductgriddatamodifier/) |  | Modify disabled product grid data | action |
| BO |  | [actionDisabledProductGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actiondisabledproductgriddefinitionmodifier/) |  | Modify disabled product grid definition | action |
| BO |  | [actionDisabledProductGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actiondisabledproductgridfilterformmodifier/) |  | Modify disabled product grid filters | action |
| BO |  | [actionDisabledProductGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actiondisabledproductgridpresentermodifier/) |  | Modify disabled product grid template data | action |
| BO |  | [actionDisabledProductGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actiondisabledproductgridquerybuildermodifier/) |  | Modify disabled product grid query builder | action |
| BO |  | [actionDiscountGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actiondiscountgriddatamodifier/) |  | Modify discount grid data | action |
| BO |  | [actionDiscountGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actiondiscountgriddefinitionmodifier/) |  | Modify discount grid definition | action |
| BO |  | [actionDiscountGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actiondiscountgridfilterformmodifier/) |  | Modify discount grid filters | action |
| BO |  | [actionDiscountGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actiondiscountgridpresentermodifier/) |  | Modify discount grid template data | action |
| BO |  | [actionDiscountGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actiondiscountgridquerybuildermodifier/) |  | Modify discount grid query builder | action |
| FO |  | [actionDispatcher](/9/modules/concepts/hooks/list-of-hooks/actiondispatcher/) |  |  | action |
| FO |  | [actionDispatcherAfter](/9/modules/concepts/hooks/list-of-hooks/actiondispatcherafter/) |  | After dispatch | action |
| FO |  | [actionDispatcherBefore](/9/modules/concepts/hooks/list-of-hooks/actiondispatcherbefore/) |  | Before dispatch | action |
| FO |  | [actionDownloadAttachment](/9/modules/concepts/hooks/list-of-hooks/actiondownloadattachment/) |  |  | action |
| FO |  | [actionDuplicateCartData](/9/modules/concepts/hooks/list-of-hooks/actionduplicatecartdata/) |  | Cart duplication | action |
| FO |  | [actionEmailAddAfterContent](/9/modules/concepts/hooks/list-of-hooks/actionemailaddaftercontent/) |  | Add extra content after mail content | action |
| FO |  | [actionEmailAddBeforeContent](/9/modules/concepts/hooks/list-of-hooks/actionemailaddbeforecontent/) |  | Add extra content before mail content | action |
| BO |  | [actionEmailLogsGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionemaillogsgriddatamodifier/) |  | Modify email logs grid data | action |
| BO |  | [actionEmailLogsGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionemaillogsgriddefinitionmodifier/) |  | Modifying E-mail grid definition | action |
| BO |  | [actionEmailLogsGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionemaillogsgridfilterformmodifier/) |  | Modify filters form for E-mail grid | action |
| BO |  | [actionEmailLogsGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionemaillogsgridpresentermodifier/) |  | Modify E-mail grid view data | action |
| BO |  | [actionEmailLogsGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionemaillogsgridquerybuildermodifier/) |  | Modify E-mail grid query builder | action |
| FO |  | [actionEmailSendBefore](/9/modules/concepts/hooks/list-of-hooks/actionemailsendbefore/) |  | Before sending an email | action |
| BO |  | [actionEmployeeGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionemployeegriddatamodifier/) |  | Modify employee grid data | action |
| BO |  | [actionEmployeeGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionemployeegriddefinitionmodifier/) |  | Modify employee grid definition | action |
| BO |  | [actionEmployeeGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionemployeegridfilterformmodifier/) |  | Modify employee grid filters | action |
| BO |  | [actionEmployeeGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionemployeegridpresentermodifier/) |  | Modify employee grid template data | action |
| BO |  | [actionEmployeeGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionemployeegridquerybuildermodifier/) |  | Modify employee grid query builder | action |
| BO |  | [actionEmptyCategoryGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionemptycategorygriddatamodifier/) |  | Modify empty category grid data | action |
| BO |  | [actionEmptyCategoryGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionemptycategorygriddefinitionmodifier/) |  | Modify empty category grid definition | action |
| BO |  | [actionEmptyCategoryGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionemptycategorygridfilterformmodifier/) |  | Modify empty category grid filters | action |
| BO |  | [actionEmptyCategoryGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionemptycategorygridpresentermodifier/) |  | Modify empty category grid template data | action |
| BO |  | [actionEmptyCategoryGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionemptycategorygridquerybuildermodifier/) |  | Modify empty category grid query builder | action |
| FO |  | [actionExportGDPRData](/9/modules/concepts/hooks/list-of-hooks/actionexportgdprdata/) |  |  | action |
| FO |  | [actionFacetedSearchCacheKeyGeneration](/9/modules/concepts/hooks/list-of-hooks/actionfacetedsearchcachekeygeneration/) |  | Customize Faceted Search cache key generation | action |
| FO |  | [actionFacetedSearchFilters](/9/modules/concepts/hooks/list-of-hooks/actionfacetedsearchfilters/) |  | Customize initial Faceted Search filters for a query | action |
| FO |  | [actionFacetedSearchSetSupportedControllers](/9/modules/concepts/hooks/list-of-hooks/actionfacetedsearchsetsupportedcontrollers/) |  | Add custom controllers to Faceted Search supported controllers list | action |
| FO |  | [actionFeatureDelete](/9/modules/concepts/hooks/list-of-hooks/actionfeaturedelete/) | afterDeleteFeature | Deleting attributes' features | action |
| FO |  | [actionFeatureSave](/9/modules/concepts/hooks/list-of-hooks/actionfeaturesave/) | afterSaveFeature | Saving attributes' features | action |
| FO |  | [actionFeatureValueDelete](/9/modules/concepts/hooks/list-of-hooks/actionfeaturevaluedelete/) | afterDeleteFeatureValue | Deleting attributes' features' values | action |
| FO |  | [actionFeatureValueSave](/9/modules/concepts/hooks/list-of-hooks/actionfeaturevaluesave/) | afterSaveFeatureValue | Saving an attributes features value | action |
| FO | ✅ | [actionFilterDeliveryOptionList](/9/modules/concepts/hooks/list-of-hooks/actionfilterdeliveryoptionlist/) |  | Modify delivery option list result | action |
| FO |  | [actionFrontControllerDetectContextCountryAfter](/9/modules/concepts/hooks/list-of-hooks/actionfrontcontrollerdetectcontextcountryafter/) |  | Action after detecting context country | action |
| FO |  | [actionFrontControllerInitAfter](/9/modules/concepts/hooks/list-of-hooks/actionfrontcontrollerinitafter/) | actionFrontControllerAfterInit | Perform actions after front office controller initialization | action |
| FO |  | [actionFrontControllerInitBefore](/9/modules/concepts/hooks/list-of-hooks/actionfrontcontrollerinitbefore/) |  | Perform actions before front office controller initialization | action |
| FO |  | [actionFrontControllerInitContextCurrencyAfter](/9/modules/concepts/hooks/list-of-hooks/actionfrontcontrollerinitcontextcurrencyafter/) |  | Action after initializing context currency | action |
| FO |  | [actionFrontControllerSetMedia](/9/modules/concepts/hooks/list-of-hooks/actionfrontcontrollersetmedia/) |  |  | action |
| FO |  | [actionFrontControllerSetVariables](/9/modules/concepts/hooks/list-of-hooks/actionfrontcontrollersetvariables/) |  | Add variables in JavaScript object and Smarty templates | action |
| FO |  | [actionFrontControllerSetVariablesBefore](/9/modules/concepts/hooks/list-of-hooks/actionfrontcontrollersetvariablesbefore/) |  | Add general purpose variables in JavaScript object and Smarty templates before assignation. | action |
| FO |  | [actionGenerateDocumentReference](/9/modules/concepts/hooks/list-of-hooks/actiongeneratedocumentreference/) |  | Modify document reference for Order | action |
| BO | ✅ | [actionGetAdminOrderButtons](/9/modules/concepts/hooks/list-of-hooks/actiongetadminorderbuttons/) |  | Admin Order Buttons | action |
| BO |  | [actionGetAdminToolbarButtons](/9/modules/concepts/hooks/list-of-hooks/actiongetadmintoolbarbuttons/) |  | Allows to add buttons in any toolbar in the back office | action |
| BO |  | [actionGetAlternativeSearchPanels](/9/modules/concepts/hooks/list-of-hooks/actiongetalternativesearchpanels/) |  | Additional search panel | action |
| FO |  | [actionGetCartRuleContextualValue](/9/modules/concepts/hooks/list-of-hooks/actiongetcartrulecontextualvalue/) |  |  | action |
| FO |  | [actionGetExtraMailTemplateVars](/9/modules/concepts/hooks/list-of-hooks/actiongetextramailtemplatevars/) |  |  | action |
| FO |  | [actionGetIDZoneByAddressID](/9/modules/concepts/hooks/list-of-hooks/actiongetidzonebyaddressid/) |  |  | action |
| FO | ✅ | [actionGetMailLayoutTransformations](/9/modules/concepts/hooks/list-of-hooks/actiongetmaillayouttransformations/) |  | Define the transformation to apply on layout | action |
| BO |  | [actionGetPdfRenderer](/9/modules/concepts/hooks/list-of-hooks/actiongetpdfrenderer/) |  | Allows modules to provide a custom PDF renderer | action |
| BO |  | [actionGetPdfTemplateObject](/9/modules/concepts/hooks/list-of-hooks/actiongetpdftemplateobject/) |  | Allows modules to provide custom PDF template objects | action |
| FO |  | [actionGetProductPropertiesAfter](/9/modules/concepts/hooks/list-of-hooks/actiongetproductpropertiesafter/) |  |  | action |
| FO |  | [actionGetProductPropertiesAfterUnitPrice](/9/modules/concepts/hooks/list-of-hooks/actiongetproductpropertiesafterunitprice/) |  | Product Properties | action |
| FO |  | [actionGetProductPropertiesBefore](/9/modules/concepts/hooks/list-of-hooks/actiongetproductpropertiesbefore/) |  |  | action |
| FO |  | [actionHtaccessCreate](/9/modules/concepts/hooks/list-of-hooks/actionhtaccesscreate/) | afterCreateHtaccess | After htaccess creation | action |
| FO |  | [actionInvoiceNumberFormatted](/9/modules/concepts/hooks/list-of-hooks/actioninvoicenumberformatted/) |  |  | action |
| BO |  | [actionLanguageGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionlanguagegriddatamodifier/) |  | Modify language grid data | action |
| BO |  | [actionLanguageGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionlanguagegriddefinitionmodifier/) |  | Modify language grid definition | action |
| BO |  | [actionLanguageGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionlanguagegridfilterformmodifier/) |  | Modify language grid filters | action |
| BO |  | [actionLanguageGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionlanguagegridpresentermodifier/) |  | Modify language grid template data | action |
| BO |  | [actionLanguageGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionlanguagegridquerybuildermodifier/) |  | Modify language grid query builder | action |
| FO |  | [actionLanguageLinkParameters](/9/modules/concepts/hooks/list-of-hooks/actionlanguagelinkparameters/) |  | Add parameters to language link | action |
| FO | ✅ | [actionListMailThemes](/9/modules/concepts/hooks/list-of-hooks/actionlistmailthemes/) |  | List the available email themes and layouts | action |
| BO | ✅ | [actionListModules](/9/modules/concepts/hooks/list-of-hooks/actionlistmodules/) |  | Allows you to add your own modules from a remote source in the modules list in the back office. | action |
| BO  FO |  | [actionLoggerLogMessage](/9/modules/concepts/hooks/list-of-hooks/actionloggerlogmessage/) |  | Allows to make extra action while a log is triggered | action |
| BO |  | [actionLogsGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionlogsgriddatamodifier/) |  | Modify logs grid data | action |
| BO |  | [actionLogsGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionlogsgriddefinitionmodifier/) |  | Modifying Logs grid definition | action |
| BO |  | [actionLogsGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionlogsgridfilterformmodifier/) |  | Modify filters form for Logs grid | action |
| BO |  | [actionLogsGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionlogsgridpresentermodifier/) |  | Modify Logs grid view data | action |
| BO |  | [actionLogsGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionlogsgridquerybuildermodifier/) |  | Modify Logs grid query builder | action |
| FO |  | [actionMailAlterMessageBeforeSend](/9/modules/concepts/hooks/list-of-hooks/actionmailaltermessagebeforesend/) |  |  | action |
| FO |  | [actionMainMenuModifier](/9/modules/concepts/hooks/list-of-hooks/actionmainmenumodifier/) |  |  | action |
| BO |  | [actionManufacturerAddressGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufactureraddressgriddatamodifier/) |  | Modify manufacturer address grid data | action |
| BO |  | [actionManufacturerAddressGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufactureraddressgriddefinitionmodifier/) |  | Modify manufacturer address grid definition | action |
| BO |  | [actionManufacturerAddressGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufactureraddressgridfilterformmodifier/) |  | Modify manufacturer address grid filters | action |
| BO |  | [actionManufacturerAddressGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufactureraddressgridpresentermodifier/) |  | Modify manufacturer address grid template data | action |
| BO |  | [actionManufacturerAddressGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufactureraddressgridquerybuildermodifier/) |  | Modify manufacturer address grid query builder | action |
| BO |  | [actionManufacturerGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufacturergriddatamodifier/) |  | Modify manufacturer grid data | action |
| BO |  | [actionManufacturerGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufacturergriddefinitionmodifier/) |  | Modify manufacturer grid definition | action |
| BO |  | [actionManufacturerGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufacturergridfilterformmodifier/) |  | Modify manufacturer grid filters | action |
| BO |  | [actionManufacturerGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufacturergridpresentermodifier/) |  | Modify manufacturer grid template data | action |
| BO |  | [actionManufacturerGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionmanufacturergridquerybuildermodifier/) |  | Modify manufacturer grid query builder | action |
| BO |  | [actionMerchandiseReturnGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionmerchandisereturngriddatamodifier/) |  | Modify merchandise return grid data | action |
| BO |  | [actionMerchandiseReturnGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionmerchandisereturngriddefinitionmodifier/) |  | Modify merchandise return grid definition | action |
| BO |  | [actionMerchandiseReturnGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionmerchandisereturngridfilterformmodifier/) |  | Modify merchandise return grid filters | action |
| BO |  | [actionMerchandiseReturnGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionmerchandisereturngridpresentermodifier/) |  | Modify merchandise return grid template data | action |
| BO |  | [actionMerchandiseReturnGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionmerchandisereturngridquerybuildermodifier/) |  | Modify merchandise return grid query builder | action |
| BO |  | [actionMetaGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionmetagriddatamodifier/) |  | Modify meta grid data | action |
| BO |  | [actionMetaGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionmetagriddefinitionmodifier/) |  | Modifying SEO and URLs grid definition | action |
| BO |  | [actionMetaGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionmetagridfilterformmodifier/) |  | Modify filters form for SEO and URLs grid | action |
| BO |  | [actionMetaGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionmetagridpresentermodifier/) |  | Modify SEO and URLs grid view data | action |
| BO |  | [actionMetaGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionmetagridquerybuildermodifier/) |  | Modify SEO and URLs grid query builder | action |
| FO | ✅ | [actionModifyFrontendSitemap](/9/modules/concepts/hooks/list-of-hooks/actionmodifyfrontendsitemap/) |  | Allows modules to add own urls (even whole new groups) to frontend sitemap. | action |
| FO |  | [actionModifyHtmlPurifierConfig](/9/modules/concepts/hooks/list-of-hooks/actionmodifyhtmlpurifierconfig/) |  | Called when configuring HTMLPurifier | action |
| FO |  | [actionModuleDisable](/9/modules/concepts/hooks/list-of-hooks/actionmoduledisable/) |  |  | action |
| FO |  | [actionModuleEnable](/9/modules/concepts/hooks/list-of-hooks/actionmoduleenable/) |  |  | action |
| FO |  | [actionModuleInstallAfter](/9/modules/concepts/hooks/list-of-hooks/actionmoduleinstallafter/) |  | actionModuleInstallAfter | action |
| FO |  | [actionModuleInstallBefore](/9/modules/concepts/hooks/list-of-hooks/actionmoduleinstallbefore/) |  | actionModuleInstallBefore | action |
| FO |  | [actionModuleMailAlertSendCustomer](/9/modules/concepts/hooks/list-of-hooks/actionmodulemailalertsendcustomer/) |  |  | action |
| FO |  | [actionModuleRegisterHookAfter](/9/modules/concepts/hooks/list-of-hooks/actionmoduleregisterhookafter/) |  |  | action |
| FO |  | [actionModuleRegisterHookBefore](/9/modules/concepts/hooks/list-of-hooks/actionmoduleregisterhookbefore/) |  |  | action |
| FO |  | [actionModuleUninstallAfter](/9/modules/concepts/hooks/list-of-hooks/actionmoduleuninstallafter/) |  | actionModuleUninstallAfter | action |
| FO |  | [actionModuleUninstallBefore](/9/modules/concepts/hooks/list-of-hooks/actionmoduleuninstallbefore/) |  | actionModuleUninstallBefore | action |
| FO |  | [actionModuleUnRegisterHookAfter](/9/modules/concepts/hooks/list-of-hooks/actionmoduleunregisterhookafter/) |  |  | action |
| FO |  | [actionModuleUnRegisterHookBefore](/9/modules/concepts/hooks/list-of-hooks/actionmoduleunregisterhookbefore/) |  |  | action |
| FO |  | [actionModuleUpgradeAfter](/9/modules/concepts/hooks/list-of-hooks/actionmoduleupgradeafter/) |  |  | action |
| FO |  | [actionNewsletterRegistrationAfter](/9/modules/concepts/hooks/list-of-hooks/actionnewsletterregistrationafter/) |  |  | action |
| FO |  | [actionNewsletterRegistrationBefore](/9/modules/concepts/hooks/list-of-hooks/actionnewsletterregistrationbefore/) |  |  | action |
| BO |  | [actionNoQtyProductWithCombinationGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithcombinationgriddatamodifier/) |  | Modify no qty product with combination grid data | action |
| BO |  | [actionNoQtyProductWithCombinationGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithcombinationgriddefinitionmodifier/) |  | Modify no qty product with combination grid definition | action |
| BO |  | [actionNoQtyProductWithCombinationGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithcombinationgridfilterformmodifier/) |  | Modify no qty product with combination grid filters | action |
| BO |  | [actionNoQtyProductWithCombinationGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithcombinationgridpresentermodifier/) |  | Modify no qty product with combination grid template data | action |
| BO |  | [actionNoQtyProductWithCombinationGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithcombinationgridquerybuildermodifier/) |  | Modify no qty product with combination grid query builder | action |
| BO |  | [actionNoQtyProductWithoutCombinationGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithoutcombinationgriddatamodifier/) |  | Modify no qty product without combination grid data | action |
| BO |  | [actionNoQtyProductWithoutCombinationGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithoutcombinationgriddefinitionmodifier/) |  | Modify no qty product without combination grid definition | action |
| BO |  | [actionNoQtyProductWithoutCombinationGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithoutcombinationgridfilterformmodifier/) |  | Modify no qty product without combination grid filters | action |
| BO |  | [actionNoQtyProductWithoutCombinationGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithoutcombinationgridpresentermodifier/) |  | Modify no qty product without combination grid template data | action |
| BO |  | [actionNoQtyProductWithoutCombinationGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionnoqtyproductwithoutcombinationgridquerybuildermodifier/) |  | Modify no qty product without combination grid query builder | action |
| BO  FO |  | [actionObject<ClassName>AddAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectclassnameaddafter/) |  |  | action |
| BO  FO |  | [actionObject<ClassName>AddBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectclassnameaddbefore/) |  |  | action |
| BO  FO |  | [actionObject<ClassName>DeleteAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectclassnamedeleteafter/) |  |  | action |
| BO  FO |  | [actionObject<ClassName>DeleteBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectclassnamedeletebefore/) |  |  | action |
| BO  FO |  | [actionObject<ClassName>UpdateAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectclassnameupdateafter/) |  |  | action |
| BO  FO |  | [actionObject<ClassName>UpdateBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectclassnameupdatebefore/) |  |  | action |
| BO  FO |  | [actionObjectAddAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectaddafter/) |  |  | action |
| BO  FO |  | [actionObjectAddBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectaddbefore/) |  |  | action |
| BO  FO |  | [actionObjectAttributeAddBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectattributeaddbefore/) |  |  | action |
| BO  FO |  | [actionObjectAttributeGroupAddBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectattributegroupaddbefore/) |  |  | action |
| BO  FO |  | [actionObjectDeleteAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectdeleteafter/) |  |  | action |
| BO  FO |  | [actionObjectDeleteBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectdeletebefore/) |  |  | action |
| BO  FO |  | [actionObjectDuplicateAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectduplicateafter/) |  | After duplicating an object | action |
| BO  FO |  | [actionObjectProductCommentValidateAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectproductcommentvalidateafter/) |  |  | action |
| BO  FO |  | [actionObjectProductInCartDeleteAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectproductincartdeleteafter/) |  | Cart product removal | action |
| BO  FO |  | [actionObjectProductInCartDeleteBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectproductincartdeletebefore/) |  | Cart product removal | action |
| BO  FO |  | [actionObjectUpdateAfter](/9/modules/concepts/hooks/list-of-hooks/actionobjectupdateafter/) |  |  | action |
| BO  FO |  | [actionObjectUpdateBefore](/9/modules/concepts/hooks/list-of-hooks/actionobjectupdatebefore/) |  |  | action |
| FO |  | [actionOnImageCutAfter](/9/modules/concepts/hooks/list-of-hooks/actiononimagecutafter/) |  |  | action |
| FO |  | [actionOnImageResizeAfter](/9/modules/concepts/hooks/list-of-hooks/actiononimageresizeafter/) |  |  | action |
| FO |  | [actionOrderEdited](/9/modules/concepts/hooks/list-of-hooks/actionorderedited/) |  | Order edited | action |
| BO |  | [actionOrderGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionordergriddatamodifier/) |  | Modify order grid data | action |
| BO |  | [actionOrderGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionordergriddefinitionmodifier/) |  | Modify order grid definition | action |
| BO |  | [actionOrderGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionordergridfilterformmodifier/) |  | Modify order grid filters | action |
| BO |  | [actionOrderGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionordergridpresentermodifier/) |  | Modify order grid template data | action |
| BO |  | [actionOrderGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionordergridquerybuildermodifier/) |  | Modify order grid query builder | action |
| FO |  | [actionOrderHasBeenDelivered](/9/modules/concepts/hooks/list-of-hooks/actionorderhasbeendelivered/) |  | Called when checking if an order has been delivered | action |
| FO |  | [actionOrderHasBeenShipped](/9/modules/concepts/hooks/list-of-hooks/actionorderhasbeenshipped/) |  | Called when checking if an order has been shipped | action |
| FO |  | [actionOrderHistoryAddAfter](/9/modules/concepts/hooks/list-of-hooks/actionorderhistoryaddafter/) |  |  | action |
| BO |  | [actionOrderMessageGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionordermessagegriddatamodifier/) |  | Modify order message grid data | action |
| BO |  | [actionOrderMessageGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionordermessagegriddefinitionmodifier/) |  | Modify order message grid definition | action |
| BO |  | [actionOrderMessageGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionordermessagegridfilterformmodifier/) |  | Modify order message grid filters | action |
| BO |  | [actionOrderMessageGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionordermessagegridpresentermodifier/) |  | Modify order message grid template data | action |
| BO |  | [actionOrderMessageGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionordermessagegridquerybuildermodifier/) |  | Modify order message grid query builder | action |
| FO |  | [actionOrderReturn](/9/modules/concepts/hooks/list-of-hooks/actionorderreturn/) | orderReturn | Returned product | action |
| BO |  | [actionOrderReturnStatesGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderreturnstatesgriddatamodifier/) |  | Modify order return states grid data | action |
| BO |  | [actionOrderReturnStatesGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderreturnstatesgriddefinitionmodifier/) |  | Modify order return states grid definition | action |
| BO |  | [actionOrderReturnStatesGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderreturnstatesgridfilterformmodifier/) |  | Modify order return states grid filters | action |
| BO |  | [actionOrderReturnStatesGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderreturnstatesgridpresentermodifier/) |  | Modify order return states grid template data | action |
| BO |  | [actionOrderReturnStatesGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderreturnstatesgridquerybuildermodifier/) |  | Modify order return states grid query builder | action |
| FO |  | [actionOrderSlipAdd](/9/modules/concepts/hooks/list-of-hooks/actionorderslipadd/) | orderSlip | Order slip creation | action |
| BO |  | [actionOrderStatesGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderstatesgriddatamodifier/) |  | Modify order states grid data | action |
| BO |  | [actionOrderStatesGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderstatesgriddefinitionmodifier/) |  | Modify order states grid definition | action |
| BO |  | [actionOrderStatesGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderstatesgridfilterformmodifier/) |  | Modify order states grid filters | action |
| BO |  | [actionOrderStatesGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderstatesgridpresentermodifier/) |  | Modify order states grid template data | action |
| BO |  | [actionOrderStatesGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionorderstatesgridquerybuildermodifier/) |  | Modify order states grid query builder | action |
| FO |  | [actionOrderStatusPostUpdate](/9/modules/concepts/hooks/list-of-hooks/actionorderstatuspostupdate/) | postUpdateOrderStatus | Post update of order status | action |
| FO |  | [actionOrderStatusUpdate](/9/modules/concepts/hooks/list-of-hooks/actionorderstatusupdate/) | updateOrderStatus | Order status update - Event | action |
| FO |  | [actionOutput<Controller>HTMLBefore](/9/modules/concepts/hooks/list-of-hooks/actionoutputcontrollerhtmlbefore/) |  |  | action |
| FO |  | [actionOutputHTMLBefore](/9/modules/concepts/hooks/list-of-hooks/actionoutputhtmlbefore/) |  | Before HTML output | action |
| BO |  | [actionOutstandingGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionoutstandinggriddatamodifier/) |  | Modify outstanding grid data | action |
| BO |  | [actionOutstandingGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionoutstandinggriddefinitionmodifier/) |  | Modify outstanding grid definition | action |
| BO |  | [actionOutstandingGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionoutstandinggridfilterformmodifier/) |  | Modify outstanding grid filters | action |
| BO |  | [actionOutstandingGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionoutstandinggridpresentermodifier/) |  | Modify outstanding grid template data | action |
| BO |  | [actionOutstandingGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionoutstandinggridquerybuildermodifier/) |  | Modify outstanding grid query builder | action |
| FO |  | [actionOverrideEmployeeImage](/9/modules/concepts/hooks/list-of-hooks/actionoverrideemployeeimage/) |  | Get Employee Image | action |
| FO |  | [actionOverrideProductQuantity](/9/modules/concepts/hooks/list-of-hooks/actionoverrideproductquantity/) |  | Override product quantity calculation | action |
| FO |  | [actionOverrideQuantityAvailableByProduct](/9/modules/concepts/hooks/list-of-hooks/actionoverridequantityavailablebyproduct/) |  | Override available quantity by product | action |
| FO |  | [actionPasswordRenew](/9/modules/concepts/hooks/list-of-hooks/actionpasswordrenew/) |  |  | action |
| FO |  | [actionPaymentCCAdd](/9/modules/concepts/hooks/list-of-hooks/actionpaymentccadd/) | paymentCCAdded | Payment CC added | action |
| FO |  | [actionPaymentConfirmation](/9/modules/concepts/hooks/list-of-hooks/actionpaymentconfirmation/) | paymentConfirm | Payment confirmation | action |
| FO |  | [actionPaymentModuleProductVarTplAfter](/9/modules/concepts/hooks/list-of-hooks/actionpaymentmoduleproductvartplafter/) |  | Triggers after product data is prepared for e-mail template | action |
| FO |  | [actionPDFInvoiceRender](/9/modules/concepts/hooks/list-of-hooks/actionpdfinvoicerender/) |  |  | action |
| FO |  | [actionPresentCart](/9/modules/concepts/hooks/list-of-hooks/actionpresentcart/) |  | Cart Presenter | action |
| FO |  | [actionPresentCategory](/9/modules/concepts/hooks/list-of-hooks/actionpresentcategory/) |  | Category Presenter | action |
| FO |  | [actionPresentManufacturer](/9/modules/concepts/hooks/list-of-hooks/actionpresentmanufacturer/) |  | Manufacturer Presenter | action |
| FO |  | [actionPresentModule](/9/modules/concepts/hooks/list-of-hooks/actionpresentmodule/) |  |  | action |
| FO |  | [actionPresentObject](/9/modules/concepts/hooks/list-of-hooks/actionpresentobject/) |  | Object Presenter | action |
| FO |  | [actionPresentOrder](/9/modules/concepts/hooks/list-of-hooks/actionpresentorder/) |  | Order Presenter | action |
| FO |  | [actionPresentOrderReturn](/9/modules/concepts/hooks/list-of-hooks/actionpresentorderreturn/) |  | Order Return Presenter | action |
| FO |  | [actionPresentPaymentOptions](/9/modules/concepts/hooks/list-of-hooks/actionpresentpaymentoptions/) |  | Payment options Presenter | action |
| FO |  | [actionPresentProduct](/9/modules/concepts/hooks/list-of-hooks/actionpresentproduct/) |  | Product Presenter | action |
| FO |  | [actionPresentProductListing](/9/modules/concepts/hooks/list-of-hooks/actionpresentproductlisting/) |  | Product Listing Presenter | action |
| FO |  | [actionPresentStore](/9/modules/concepts/hooks/list-of-hooks/actionpresentstore/) |  | Store Presenter | action |
| FO |  | [actionPresentSupplier](/9/modules/concepts/hooks/list-of-hooks/actionpresentsupplier/) |  | Supplier Presenter | action |
| BO |  | [actionProductActivation](/9/modules/concepts/hooks/list-of-hooks/actionproductactivation/) |  |  | action |
| FO |  | [actionProductAdd](/9/modules/concepts/hooks/list-of-hooks/actionproductadd/) | addproduct | Product creation | action |
| FO |  | [actionProductAttributeDelete](/9/modules/concepts/hooks/list-of-hooks/actionproductattributedelete/) | deleteProductAttribute | Product attribute deletion | action |
| FO |  | [actionProductAttributeUpdate](/9/modules/concepts/hooks/list-of-hooks/actionproductattributeupdate/) | updateProductAttribute | Product attribute update | action |
| FO |  | [actionProductCancel](/9/modules/concepts/hooks/list-of-hooks/actionproductcancel/) | cancelProduct | Product cancelled | action |
| FO |  | [actionProductCoverage](/9/modules/concepts/hooks/list-of-hooks/actionproductcoverage/) |  |  | action |
| FO |  | [actionProductDelete](/9/modules/concepts/hooks/list-of-hooks/actionproductdelete/) | deleteproduct | Product deletion | action |
| FO |  | [actionProductGetAttributesGroupsAfter](/9/modules/concepts/hooks/list-of-hooks/actionproductgetattributesgroupsafter/) |  | Triggers after getting product attributes groups | action |
| FO |  | [actionProductGetAttributesGroupsBefore](/9/modules/concepts/hooks/list-of-hooks/actionproductgetattributesgroupsbefore/) |  | Triggers before getting product attributes groups | action |
| BO |  | [actionProductGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductgriddatamodifier/) |  | Modify product grid data | action |
| BO |  | [actionProductGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductgriddefinitionmodifier/) |  | Modify product grid definition | action |
| BO |  | [actionProductGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductgridfilterformmodifier/) |  | Modify product grid filters | action |
| BO |  | [actionProductGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductgridpresentermodifier/) |  | Modify product grid template data | action |
| BO |  | [actionProductGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductgridquerybuildermodifier/) |  | Modify product grid query builder | action |
| FO |  | [actionProductOutOfStock](/9/modules/concepts/hooks/list-of-hooks/actionproductoutofstock/) | productOutOfStock | Out-of-stock product | action |
| BO |  | [actionProductPriceCalculation](/9/modules/concepts/hooks/list-of-hooks/actionproductpricecalculation/) |  | Product Price Calculation | action |
| FO |  | [actionProductSave](/9/modules/concepts/hooks/list-of-hooks/actionproductsave/) | afterSaveProduct | Saving products | action |
| FO |  | [actionProductSearchAfter](/9/modules/concepts/hooks/list-of-hooks/actionproductsearchafter/) |  | Event triggered after search product completed | action |
| FO |  | [actionProductSearchProviderRunQueryAfter](/9/modules/concepts/hooks/list-of-hooks/actionproductsearchproviderrunqueryafter/) |  | Runs an action after ProductSearchProviderInterface::RunQuery() | action |
| FO |  | [actionProductSearchProviderRunQueryBefore](/9/modules/concepts/hooks/list-of-hooks/actionproductsearchproviderrunquerybefore/) |  | Runs an action before ProductSearchProviderInterface::RunQuery() | action |
| BO |  | [actionProductUpdate](/9/modules/concepts/hooks/list-of-hooks/actionproductupdate/) | updateproduct | Product update | action |
| BO |  | [actionProductWithoutDescriptionGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutdescriptiongriddatamodifier/) |  | Modify product without description grid data | action |
| BO |  | [actionProductWithoutDescriptionGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutdescriptiongriddefinitionmodifier/) |  | Modify product without description grid definition | action |
| BO |  | [actionProductWithoutDescriptionGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutdescriptiongridfilterformmodifier/) |  | Modify product without description grid filters | action |
| BO |  | [actionProductWithoutDescriptionGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutdescriptiongridpresentermodifier/) |  | Modify product without description grid template data | action |
| BO |  | [actionProductWithoutDescriptionGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutdescriptiongridquerybuildermodifier/) |  | Modify product without description grid query builder | action |
| BO |  | [actionProductWithoutImageGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutimagegriddatamodifier/) |  | Modify product without image grid data | action |
| BO |  | [actionProductWithoutImageGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutimagegriddefinitionmodifier/) |  | Modify product without image grid definition | action |
| BO |  | [actionProductWithoutImageGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutimagegridfilterformmodifier/) |  | Modify product without image grid filters | action |
| BO |  | [actionProductWithoutImageGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutimagegridpresentermodifier/) |  | Modify product without image grid template data | action |
| BO |  | [actionProductWithoutImageGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutimagegridquerybuildermodifier/) |  | Modify product without image grid query builder | action |
| BO |  | [actionProductWithoutPriceGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutpricegriddatamodifier/) |  | Modify product without price grid data | action |
| BO |  | [actionProductWithoutPriceGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutpricegriddefinitionmodifier/) |  | Modify product without price grid definition | action |
| BO |  | [actionProductWithoutPriceGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutpricegridfilterformmodifier/) |  | Modify product without price grid filters | action |
| BO |  | [actionProductWithoutPriceGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutpricegridpresentermodifier/) |  | Modify product without price grid template data | action |
| BO |  | [actionProductWithoutPriceGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionproductwithoutpricegridquerybuildermodifier/) |  | Modify product without price grid query builder | action |
| BO |  | [actionProfileGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionprofilegriddatamodifier/) |  | Modify profile grid data | action |
| BO |  | [actionProfileGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionprofilegriddefinitionmodifier/) |  | Modify profile grid definition | action |
| BO |  | [actionProfileGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionprofilegridfilterformmodifier/) |  | Modify profile grid filters | action |
| BO |  | [actionProfileGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionprofilegridpresentermodifier/) |  | Modify profile grid template data | action |
| BO |  | [actionProfileGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionprofilegridquerybuildermodifier/) |  | Modify profile grid query builder | action |
| FO |  | [actionSearch](/9/modules/concepts/hooks/list-of-hooks/actionsearch/) |  |  | action |
| BO |  | [actionSearchEngineGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionsearchenginegriddatamodifier/) |  | Modify search engine grid data | action |
| BO |  | [actionSearchEngineGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionsearchenginegriddefinitionmodifier/) |  | Modify search engine grid definition | action |
| BO |  | [actionSearchEngineGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionsearchenginegridfilterformmodifier/) |  | Modify search engine grid filters | action |
| BO |  | [actionSearchEngineGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionsearchenginegridpresentermodifier/) |  | Modify search engine grid template data | action |
| BO |  | [actionSearchEngineGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionsearchenginegridquerybuildermodifier/) |  | Modify search engine grid query builder | action |
| BO |  | [actionSecuritySessionCustomerGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessioncustomergriddatamodifier/) |  | Modify security session customer grid data | action |
| BO |  | [actionSecuritySessionCustomerGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessioncustomergriddefinitionmodifier/) |  | Modify security session customer grid definition | action |
| BO |  | [actionSecuritySessionCustomerGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessioncustomergridfilterformmodifier/) |  | Modify security session customer grid filters | action |
| BO |  | [actionSecuritySessionCustomerGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessioncustomergridpresentermodifier/) |  | Modify security session customer grid template data | action |
| BO |  | [actionSecuritySessionCustomerGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessioncustomergridquerybuildermodifier/) |  | Modify security session customer grid query builder | action |
| BO |  | [actionSecuritySessionEmployeeGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessionemployeegriddatamodifier/) |  | Modify security session employee grid data | action |
| BO |  | [actionSecuritySessionEmployeeGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessionemployeegriddefinitionmodifier/) |  | Modify security session employee grid definition | action |
| BO |  | [actionSecuritySessionEmployeeGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessionemployeegridfilterformmodifier/) |  | Modify security session employee grid filters | action |
| BO |  | [actionSecuritySessionEmployeeGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessionemployeegridpresentermodifier/) |  | Modify security session employee grid template data | action |
| BO |  | [actionSecuritySessionEmployeeGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionsecuritysessionemployeegridquerybuildermodifier/) |  | Modify security session employee grid query builder | action |
| FO |  | [actionSetInvoice](/9/modules/concepts/hooks/list-of-hooks/actionsetinvoice/) |  |  | action |
| FO |  | [actionShopDataDuplication](/9/modules/concepts/hooks/list-of-hooks/actionshopdataduplication/) |  |  | action |
| BO |  | [actionSqlRequestGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionsqlrequestgriddatamodifier/) |  | Modify sql request grid data | action |
| BO |  | [actionSqlRequestGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionsqlrequestgriddefinitionmodifier/) |  | Modifying SQL Manager grid definition | action |
| BO |  | [actionSqlRequestGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionsqlrequestgridfilterformmodifier/) |  | Modify filters form for SQL Manager grid | action |
| BO |  | [actionSqlRequestGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionsqlrequestgridpresentermodifier/) |  | Modify SQL Manager grid view data | action |
| BO |  | [actionSqlRequestGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionsqlrequestgridquerybuildermodifier/) |  | Modify SQL Manager grid query builder | action |
| BO |  | [actionStateGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionstategriddatamodifier/) |  | Modify state grid data | action |
| BO |  | [actionStateGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionstategriddefinitionmodifier/) |  | Modify state grid definition | action |
| BO |  | [actionStateGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionstategridfilterformmodifier/) |  | Modify state grid filters | action |
| BO |  | [actionStateGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionstategridpresentermodifier/) |  | Modify state grid template data | action |
| BO |  | [actionStateGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionstategridquerybuildermodifier/) |  | Modify state grid query builder | action |
| FO |  | [actionSubmitAccountBefore](/9/modules/concepts/hooks/list-of-hooks/actionsubmitaccountbefore/) | actionBeforeSubmitAccount |  | action |
| FO |  | [actionSubmitCustomerAddressForm](/9/modules/concepts/hooks/list-of-hooks/actionsubmitcustomeraddressform/) |  |  | action |
| BO |  | [actionSupplierGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionsuppliergriddatamodifier/) |  | Modify supplier grid data | action |
| BO |  | [actionSupplierGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionsuppliergriddefinitionmodifier/) |  | Modify supplier grid definition | action |
| BO |  | [actionSupplierGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionsuppliergridfilterformmodifier/) |  | Modify supplier grid filters | action |
| BO |  | [actionSupplierGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionsuppliergridpresentermodifier/) |  | Modify supplier grid template data | action |
| BO |  | [actionSupplierGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionsuppliergridquerybuildermodifier/) |  | Modify supplier grid query builder | action |
| BO |  | [actionTaxGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxgriddatamodifier/) |  | Modify tax grid data | action |
| BO |  | [actionTaxGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxgriddefinitionmodifier/) |  | Modify tax grid definition | action |
| BO |  | [actionTaxGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxgridfilterformmodifier/) |  | Modify tax grid filters | action |
| BO |  | [actionTaxGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxgridpresentermodifier/) |  | Modify tax grid template data | action |
| BO |  | [actionTaxGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxgridquerybuildermodifier/) |  | Modify tax grid query builder | action |
| FO |  | [actionTaxManager](/9/modules/concepts/hooks/list-of-hooks/actiontaxmanager/) | 116, 97, 120, 77, 97, 110, 97, 103, 101, 114 | Tax Manager Factory | action |
| BO |  | [actionTaxRulesGroupGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxrulesgroupgriddatamodifier/) |  | Modify tax rules group grid data | action |
| BO |  | [actionTaxRulesGroupGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxrulesgroupgriddefinitionmodifier/) |  | Modify tax rules group grid definition | action |
| BO |  | [actionTaxRulesGroupGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxrulesgroupgridfilterformmodifier/) |  | Modify tax rules group grid filters | action |
| BO |  | [actionTaxRulesGroupGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxrulesgroupgridpresentermodifier/) |  | Modify tax rules group grid template data | action |
| BO |  | [actionTaxRulesGroupGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actiontaxrulesgroupgridquerybuildermodifier/) |  | Modify tax rules group grid query builder | action |
| BO |  | [actionTitleGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actiontitlegriddatamodifier/) |  | Modify title grid data | action |
| BO |  | [actionTitleGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actiontitlegriddefinitionmodifier/) |  | Modify title grid definition | action |
| BO |  | [actionTitleGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actiontitlegridfilterformmodifier/) |  | Modify title grid filters | action |
| BO |  | [actionTitleGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actiontitlegridpresentermodifier/) |  | Modify title grid template data | action |
| BO |  | [actionTitleGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actiontitlegridquerybuildermodifier/) |  | Modify title grid query builder | action |
| FO |  | [actionUpdateCartAddress](/9/modules/concepts/hooks/list-of-hooks/actionupdatecartaddress/) |  | Triggers after changing address on the cart | action |
| BO |  | [actionUpdateDefaultCombinationAfter](/9/modules/concepts/hooks/list-of-hooks/actionupdatedefaultcombinationafter/) |  | React after a product default combination is updated | action |
| FO |  | [actionUpdateLangAfter](/9/modules/concepts/hooks/list-of-hooks/actionupdatelangafter/) |  | Update "lang" tables | action |
| FO |  | [actionUpdateQuantity](/9/modules/concepts/hooks/list-of-hooks/actionupdatequantity/) | updateQuantity | Quantity update | action |
| FO |  | [actionValidateCartRule](/9/modules/concepts/hooks/list-of-hooks/actionvalidatecartrule/) |  | Alter cart rule validation process | action |
| FO |  | [actionValidateCustomerAddressForm](/9/modules/concepts/hooks/list-of-hooks/actionvalidatecustomeraddressform/) |  | Customer address form validation | action |
| FO | ✅ | [actionValidateOrder](/9/modules/concepts/hooks/list-of-hooks/actionvalidateorder/) | newOrder | New orders | action |
| FO |  | [actionValidateOrderAfter](/9/modules/concepts/hooks/list-of-hooks/actionvalidateorderafter/) |  | After validating an order | action |
| FO |  | [actionValidateOrderBefore](/9/modules/concepts/hooks/list-of-hooks/actionvalidateorderbefore/) |  | Before validating an order | action |
| FO |  | [actionValidateStepComplete](/9/modules/concepts/hooks/list-of-hooks/actionvalidatestepcomplete/) |  |  | action |
| FO |  | [actionWatermark](/9/modules/concepts/hooks/list-of-hooks/actionwatermark/) | watermark | Watermark | action |
| BO |  | [actionWebserviceKeyGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionwebservicekeygriddatamodifier/) |  | Modify webservice key grid data | action |
| BO |  | [actionWebserviceKeyGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionwebservicekeygriddefinitionmodifier/) |  | Modifying Webservice grid definition | action |
| BO |  | [actionWebserviceKeyGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionwebservicekeygridfilterformmodifier/) |  | Modify filters form for Webservice grid | action |
| BO |  | [actionWebserviceKeyGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionwebservicekeygridpresentermodifier/) |  | Modify Webservice grid view data | action |
| BO |  | [actionWebserviceKeyGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionwebservicekeygridquerybuildermodifier/) |  | Modify Webservice grid query builder | action |
| FO |  | [actionWishlistAddProduct](/9/modules/concepts/hooks/list-of-hooks/actionwishlistaddproduct/) |  |  | action |
| BO |  | [actionZoneGridDataModifier](/9/modules/concepts/hooks/list-of-hooks/actionzonegriddatamodifier/) |  | Modify zone grid data | action |
| BO |  | [actionZoneGridDefinitionModifier](/9/modules/concepts/hooks/list-of-hooks/actionzonegriddefinitionmodifier/) |  | Modify zone grid definition | action |
| BO |  | [actionZoneGridFilterFormModifier](/9/modules/concepts/hooks/list-of-hooks/actionzonegridfilterformmodifier/) |  | Modify zone grid filters | action |
| BO |  | [actionZoneGridPresenterModifier](/9/modules/concepts/hooks/list-of-hooks/actionzonegridpresentermodifier/) |  | Modify zone grid template data | action |
| BO |  | [actionZoneGridQueryBuilderModifier](/9/modules/concepts/hooks/list-of-hooks/actionzonegridquerybuildermodifier/) |  | Modify zone grid query builder | action |
| FO |  | [adaptImageLink](/9/modules/concepts/hooks/list-of-hooks/adaptimagelink/) |  | Adapt product image link | action |
| FO |  | [additionalCustomerAddressFields](/9/modules/concepts/hooks/list-of-hooks/additionalcustomeraddressfields/) |  | Add fields to the Customer address form |  |
| FO |  | [additionalCustomerFormFields](/9/modules/concepts/hooks/list-of-hooks/additionalcustomerformfields/) |  | Add fields to the Customer form |  |
| FO |  | [additionalHtmlAttributesFormFields](/9/modules/concepts/hooks/list-of-hooks/additionalhtmlattributesformfields/) |  |  | action |
| FO |  | [addWebserviceResources](/9/modules/concepts/hooks/list-of-hooks/addwebserviceresources/) |  | Add extra webservice resource |  |
| BO |  | [dashboardData](/9/modules/concepts/hooks/list-of-hooks/dashboarddata/) |  |  |  |
| BO |  | [dashboardZoneOne](/9/modules/concepts/hooks/list-of-hooks/dashboardzoneone/) |  |  |  |
| BO |  | [dashboardZoneThree](/9/modules/concepts/hooks/list-of-hooks/dashboardzonethree/) |  | Dashboard column three |  |
| BO |  | [dashboardZoneTwo](/9/modules/concepts/hooks/list-of-hooks/dashboardzonetwo/) |  |  |  |
| FO |  | [deleteProductAttribute](/9/modules/concepts/hooks/list-of-hooks/deleteproductattribute/) |  |  |  |
| FO |  | [display<Controller>Header](/9/modules/concepts/hooks/list-of-hooks/displaycontrollerheader/) |  |  | display |
| FO |  | [displayAdditionalCustomerAddressFields](/9/modules/concepts/hooks/list-of-hooks/displayadditionalcustomeraddressfields/) |  | Display additional customer address fields | display |
| FO |  | [displayAddressSelectorBottom](/9/modules/concepts/hooks/list-of-hooks/displayaddressselectorbottom/) |  |  | display |
| BO |  | [displayAdminAfterHeader](/9/modules/concepts/hooks/list-of-hooks/displayadminafterheader/) |  |  | display |
| BO |  | [displayAdminCustomers](/9/modules/concepts/hooks/list-of-hooks/displayadmincustomers/) | adminCustomers | Display new elements in the Back Office, tab AdminCustomers | display |
| BO |  | [displayAdminEndContent](/9/modules/concepts/hooks/list-of-hooks/displayadminendcontent/) |  | Administration end of content | display |
| BO |  | [displayAdminForm](/9/modules/concepts/hooks/list-of-hooks/displayadminform/) |  |  | display |
| BO |  | [displayAdminGridTableAfter](/9/modules/concepts/hooks/list-of-hooks/displayadmingridtableafter/) | displayAdminListAfter | Display after Grid table | display |
| BO |  | [displayAdminGridTableBefore](/9/modules/concepts/hooks/list-of-hooks/displayadmingridtablebefore/) | displayAdminListBefore | Display before Grid table | display |
| BO |  | [displayAdminListAfter](/9/modules/concepts/hooks/list-of-hooks/displayadminlistafter/) |  |  | display |
| BO |  | [displayAdminListBefore](/9/modules/concepts/hooks/list-of-hooks/displayadminlistbefore/) |  |  | display |
| BO |  | [displayAdminLogin](/9/modules/concepts/hooks/list-of-hooks/displayadminlogin/) |  |  | display |
| BO |  | [displayAdminNavBarBeforeEnd](/9/modules/concepts/hooks/list-of-hooks/displayadminnavbarbeforeend/) |  |  | display |
| BO |  | [displayAdminOptions](/9/modules/concepts/hooks/list-of-hooks/displayadminoptions/) |  |  | display |
| BO | ✅ | [displayAdminOrder](/9/modules/concepts/hooks/list-of-hooks/displayadminorder/) | adminOrder | Display new elements in the Back Office, tab AdminOrder | display |
| BO |  | [displayAdminOrderCreateExtraButtons](/9/modules/concepts/hooks/list-of-hooks/displayadminordercreateextrabuttons/) |  | Add buttons on the create order page dropdown | display |
| BO | ✅ | [displayAdminOrderMain](/9/modules/concepts/hooks/list-of-hooks/displayadminordermain/) |  | Admin Order Main Column | display |
| BO |  | [displayAdminOrderMainBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminordermainbottom/) |  | Admin Order Main Column Bottom | display |
| BO |  | [displayAdminOrderSide](/9/modules/concepts/hooks/list-of-hooks/displayadminorderside/) | displayBackofficeOrderActions | Admin Order Side Column | display |
| BO | ✅ | [displayAdminOrderSideBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminordersidebottom/) |  |  | display |
| BO | ✅ | [displayAdminOrderTabContent](/9/modules/concepts/hooks/list-of-hooks/displayadminordertabcontent/) |  | Display new elements in the Back Office, tab contents on order | display |
| BO | ✅ | [displayAdminOrderTabLink](/9/modules/concepts/hooks/list-of-hooks/displayadminordertablink/) |  | Display new elements in the Back Office, tab links on order | display |
| BO | ✅ | [displayAdminOrderTop](/9/modules/concepts/hooks/list-of-hooks/displayadminordertop/) |  | Display new elements in the Back Office, top of Order page | display |
| BO |  | [displayAdminProductsCombinationBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductscombinationbottom/) |  | Display new elements in back office product page, Combination tab | display |
| BO | ✅ | [displayAdminProductsExtra](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsextra/) |  | Display new elements in back office product page, Extra tab | display |
| BO |  | [displayAdminProductsMainStepLeftColumnBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsmainstepleftcolumnbottom/) |  | Display new elements in back office product page, left column of the Basic settings tab | display |
| BO |  | [displayAdminProductsMainStepLeftColumnMiddle](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsmainstepleftcolumnmiddle/) |  | Display new elements in back office product page, left column of the Basic settings tab | display |
| BO |  | [displayAdminProductsMainStepRightColumnBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsmainsteprightcolumnbottom/) |  | Display new elements in back office product page, right column of the Basic settings tab | display |
| BO |  | [displayAdminProductsOptionsStepBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsoptionsstepbottom/) |  | Display new elements in back office product page, Options tab | display |
| BO |  | [displayAdminProductsOptionsStepTop](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsoptionssteptop/) |  | Display new elements in back office product page, Options tab | display |
| BO |  | [displayAdminProductsPriceStepBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductspricestepbottom/) |  | Display new elements in back office product page, Price tab | display |
| BO | ✅ | [displayAdminProductsQuantitiesStepBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsquantitiesstepbottom/) |  | Display new elements in back office product page, Quantities/Combinations tab | display |
| BO |  | [displayAdminProductsSeoStepBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsseostepbottom/) |  | Display new elements in back office product page, SEO tab | display |
| BO |  | [displayAdminProductsShippingStepBottom](/9/modules/concepts/hooks/list-of-hooks/displayadminproductsshippingstepbottom/) |  | Display new elements in back office product page, Shipping tab | display |
| BO |  | [displayAdminStatsGraphEngine](/9/modules/concepts/hooks/list-of-hooks/displayadminstatsgraphengine/) | GraphEngine |  | display |
| BO |  | [displayAdminStatsGridEngine](/9/modules/concepts/hooks/list-of-hooks/displayadminstatsgridengine/) | GridEngine |  | display |
| BO |  | [displayAdminStatsModules](/9/modules/concepts/hooks/list-of-hooks/displayadminstatsmodules/) | AdminStatsModules | Stats - Modules | display |
| BO |  | [displayAdminStoreInformation](/9/modules/concepts/hooks/list-of-hooks/displayadminstoreinformation/) |  | Display extra store information | display |
| BO |  | [displayAdminThemesListAfter](/9/modules/concepts/hooks/list-of-hooks/displayadminthemeslistafter/) |  | BO themes list extra content | display |
| BO |  | [displayAdminView](/9/modules/concepts/hooks/list-of-hooks/displayadminview/) |  |  | display |
| FO |  | [displayAfterBodyOpeningTag](/9/modules/concepts/hooks/list-of-hooks/displayafterbodyopeningtag/) |  | Very top of pages | display |
| FO |  | [displayAfterCarrier](/9/modules/concepts/hooks/list-of-hooks/displayaftercarrier/) |  | After carriers list | display |
| FO |  | [displayAfterProductThumbs](/9/modules/concepts/hooks/list-of-hooks/displayafterproductthumbs/) |  | Display extra content below product thumbs | display |
| FO |  | [displayAfterTitleTag](/9/modules/concepts/hooks/list-of-hooks/displayaftertitletag/) |  | After title tag | display |
| BO |  | [displayAttributeForm](/9/modules/concepts/hooks/list-of-hooks/displayattributeform/) | attributeForm |  | display |
| BO |  | [displaybackOfficeEmployeeMenu](/9/modules/concepts/hooks/list-of-hooks/displaybackofficeemployeemenu/) |  | Administration Employee menu | display |
| BO |  | [displayBackOfficeFooter](/9/modules/concepts/hooks/list-of-hooks/displaybackofficefooter/) | 100, 105, 115, 112, 108, 97, 121, 66, 97, 99, 107, 79, 102, 102, 105, 99, 101, 70, 111, 111, 116, 101, 114 | Administration panel footer | display |
| BO |  | [displaybackOfficeHeader](/9/modules/concepts/hooks/list-of-hooks/displaybackofficeheader/) | backOfficeHeader | Administration panel header | display |
| BO |  | [displaybackOfficeTop](/9/modules/concepts/hooks/list-of-hooks/displaybackofficetop/) | backOfficeTop | Administration panel hover the tabs | display |
| FO |  | [displayBanner](/9/modules/concepts/hooks/list-of-hooks/displaybanner/) |  | Very top of pages | display |
| FO |  | [displayBeforeBodyClosingTag](/9/modules/concepts/hooks/list-of-hooks/displaybeforebodyclosingtag/) |  | Very bottom of pages | display |
| FO |  | [displayBeforeCarrier](/9/modules/concepts/hooks/list-of-hooks/displaybeforecarrier/) | beforeCarrier | Before carriers list | display |
| FO |  | [displayCarrierExtraContent](/9/modules/concepts/hooks/list-of-hooks/displaycarrierextracontent/) |  | Display additional content for a carrier (e.g pickup points) | display |
| FO |  | [displayCartExtraProductActions](/9/modules/concepts/hooks/list-of-hooks/displaycartextraproductactions/) |  | Extra buttons in shopping cart | action |
| FO |  | [displayCartExtraProductInfo](/9/modules/concepts/hooks/list-of-hooks/displaycartextraproductinfo/) |  | Extra information in shopping cart product line | display |
| FO |  | [displayCartModalContent](/9/modules/concepts/hooks/list-of-hooks/displaycartmodalcontent/) |  | Content of Add-to-cart modal | display |
| FO |  | [displayCartModalFooter](/9/modules/concepts/hooks/list-of-hooks/displaycartmodalfooter/) |  | Bottom of Add-to-cart modal | display |
| FO |  | [displayCheckoutBeforeConfirmation](/9/modules/concepts/hooks/list-of-hooks/displaycheckoutbeforeconfirmation/) |  | Show custom content before checkout confirmation | display |
| FO |  | [displayCheckoutSubtotalDetails](/9/modules/concepts/hooks/list-of-hooks/displaycheckoutsubtotaldetails/) |  | displayCheckoutSubtotalDetails | display |
| FO |  | [displayCheckoutSummaryTop](/9/modules/concepts/hooks/list-of-hooks/displaycheckoutsummarytop/) |  | Cart summary top | display |
| FO |  | [displayCMSDisputeInformation](/9/modules/concepts/hooks/list-of-hooks/displaycmsdisputeinformation/) |  | displayCMSDisputeInformation | display |
| FO |  | [displayCMSPrintButton](/9/modules/concepts/hooks/list-of-hooks/displaycmsprintbutton/) |  | displayCMSPrintButton | display |
| FO |  | [displayContactContent](/9/modules/concepts/hooks/list-of-hooks/displaycontactcontent/) |  | Content wrapper section of the contact page | display |
| FO |  | [displayContactLeftColumn](/9/modules/concepts/hooks/list-of-hooks/displaycontactleftcolumn/) |  | Left column blocks on the contact page | display |
| FO |  | [displayContactRightColumn](/9/modules/concepts/hooks/list-of-hooks/displaycontactrightcolumn/) |  | Right column blocks of the contact page | display |
| FO |  | [displayContentWrapperBottom](/9/modules/concepts/hooks/list-of-hooks/displaycontentwrapperbottom/) |  | Content wrapper section (bottom) | display |
| FO |  | [displayContentWrapperTop](/9/modules/concepts/hooks/list-of-hooks/displaycontentwrappertop/) |  | Content wrapper section (top) | display |
| FO |  | [displayCrossSellingShoppingCart](/9/modules/concepts/hooks/list-of-hooks/displaycrosssellingshoppingcart/) |  | displayCrossSellingShoppingCart | display |
| FO |  | [displayCustomerAccount](/9/modules/concepts/hooks/list-of-hooks/displaycustomeraccount/) | customerAccount | Customer account displayed in Front Office | display |
| FO |  | [displayCustomerAccountForm](/9/modules/concepts/hooks/list-of-hooks/displaycustomeraccountform/) | createAccountForm | Customer account creation form | display |
| FO |  | [displayCustomerAccountFormTop](/9/modules/concepts/hooks/list-of-hooks/displaycustomeraccountformtop/) | createAccountTop | Block above the form for create an account | display |
| FO |  | [displayCustomerAccountTop](/9/modules/concepts/hooks/list-of-hooks/displaycustomeraccounttop/) |  | Customer account displayed in Front Office (Top part) | display |
| FO |  | [displayCustomerLoginFormAfter](/9/modules/concepts/hooks/list-of-hooks/displaycustomerloginformafter/) |  | Display elements after login form | display |
| FO |  | [displayCustomization](/9/modules/concepts/hooks/list-of-hooks/displaycustomization/) |  |  | display |
| BO |  | [displayDashboardToolbarIcons](/9/modules/concepts/hooks/list-of-hooks/displaydashboardtoolbaricons/) |  | Display new elements in back office page with dashboard, on icons list | display |
| BO |  | [displayDashboardToolbarTopMenu](/9/modules/concepts/hooks/list-of-hooks/displaydashboardtoolbartopmenu/) |  | Display new elements in back office page with a dashboard, on top Menu | display |
| BO |  | [displayDashboardTop](/9/modules/concepts/hooks/list-of-hooks/displaydashboardtop/) |  | Dashboard Top | display |
| BO |  | [displayEmptyModuleCategoryExtraMessage](/9/modules/concepts/hooks/list-of-hooks/displayemptymodulecategoryextramessage/) |  | Extra message to display for an empty modules category | display |
| FO |  | [displayExpressCheckout](/9/modules/concepts/hooks/list-of-hooks/displayexpresscheckout/) |  | displayExpressCheckout | display |
| BO |  | [displayFeatureForm](/9/modules/concepts/hooks/list-of-hooks/displayfeatureform/) | featureForm | Add fields to the form 'feature' | display |
| BO |  | [displayFeaturePostProcess](/9/modules/concepts/hooks/list-of-hooks/displayfeaturepostprocess/) | postProcessFeature | On post-process in admin feature | display |
| BO |  | [displayFeatureValueForm](/9/modules/concepts/hooks/list-of-hooks/displayfeaturevalueform/) | featureValueForm | Add fields to the form 'feature value' | display |
| BO |  | [displayFeatureValuePostProcess](/9/modules/concepts/hooks/list-of-hooks/displayfeaturevaluepostprocess/) | postProcessFeatureValue | On post-process in admin feature value | display |
| FO |  | [displayFooter](/9/modules/concepts/hooks/list-of-hooks/displayfooter/) | footer | Footer | display |
| FO |  | [displayFooterAfter](/9/modules/concepts/hooks/list-of-hooks/displayfooterafter/) |  | displayFooterAfter | display |
| FO |  | [displayFooterBefore](/9/modules/concepts/hooks/list-of-hooks/displayfooterbefore/) |  | displayFooterBefore | display |
| FO |  | [displayFooterCategory](/9/modules/concepts/hooks/list-of-hooks/displayfootercategory/) |  | This hook adds new blocks under the products listing in a category/search | display |
| FO |  | [displayFooterProduct](/9/modules/concepts/hooks/list-of-hooks/displayfooterproduct/) | productfooter | Product footer | display |
| FO |  | [displayGDPRConsent](/9/modules/concepts/hooks/list-of-hooks/displaygdprconsent/) |  | displayGDPRConsent | display |
| FO |  | [displayHeader](/9/modules/concepts/hooks/list-of-hooks/displayheader/) | Header | Pages html head section | display |
| FO |  | [displayHeaderCategory](/9/modules/concepts/hooks/list-of-hooks/displayheadercategory/) |  | Footer above of the product list | display |
| FO |  | [displayHome](/9/modules/concepts/hooks/list-of-hooks/displayhome/) | home | Homepage content | display |
| FO |  | [displayInvoiceLegalFreeText](/9/modules/concepts/hooks/list-of-hooks/displayinvoicelegalfreetext/) |  | PDF Invoice - Legal Free Text | display |
| FO |  | [displayLeftColumn](/9/modules/concepts/hooks/list-of-hooks/displayleftcolumn/) | extraLeft | New elements on a page (left column) | display |
| FO |  | [displayLeftColumnProduct](/9/modules/concepts/hooks/list-of-hooks/displayleftcolumnproduct/) | extraLeft | New elements on the product page (left column) | display |
| FO |  | [displayMaintenance](/9/modules/concepts/hooks/list-of-hooks/displaymaintenance/) |  | Maintenance Page | display |
| BO |  | [displayModuleConfigureExtraButons](/9/modules/concepts/hooks/list-of-hooks/displaymoduleconfigureextrabuttons/) |  | Module configuration - After toolbar buttons | display |
| FO |  | [displayMyAccountBlock](/9/modules/concepts/hooks/list-of-hooks/displaymyaccountblock/) | myAccountBlock | My account block | display |
| FO |  | [displayNav1](/9/modules/concepts/hooks/list-of-hooks/displaynav1/) |  | displayNav1 | display |
| FO |  | [displayNav2](/9/modules/concepts/hooks/list-of-hooks/displaynav2/) |  | displayNav2 | display |
| FO |  | [displayNavFullWidth](/9/modules/concepts/hooks/list-of-hooks/displaynavfullwidth/) |  | Navigation | display |
| FO |  | [displayNewsletterRegistration](/9/modules/concepts/hooks/list-of-hooks/displaynewsletterregistration/) |  | displayNewsletterRegistration | display |
| FO |  | [displayNotFound](/9/modules/concepts/hooks/list-of-hooks/displaynotfound/) |  | displayNotFound | display |
| FO |  | [displayOrderConfirmation](/9/modules/concepts/hooks/list-of-hooks/displayorderconfirmation/) | orderConfirmation | Order confirmation page | display |
| FO |  | [displayOrderConfirmation1](/9/modules/concepts/hooks/list-of-hooks/displayorderconfirmation1/) |  | displayOrderConfirmation1 | display |
| FO |  | [displayOrderConfirmation2](/9/modules/concepts/hooks/list-of-hooks/displayorderconfirmation2/) |  | displayOrderConfirmation2 | display |
| FO |  | [displayOrderDetail](/9/modules/concepts/hooks/list-of-hooks/displayorderdetail/) | orderDetailDisplayed | Order detail | display |
| BO |  | [displayOrderPreview](/9/modules/concepts/hooks/list-of-hooks/displayorderpreview/) |  |  | display |
| FO |  | [displayOverrideTemplate](/9/modules/concepts/hooks/list-of-hooks/displayoverridetemplate/) |  | Change the default template of current controller | display |
| FO |  | [displayPaymentByBinaries](/9/modules/concepts/hooks/list-of-hooks/displaypaymentbybinaries/) |  | Payment form generated by binaries | display |
| FO |  | [displayPaymentReturn](/9/modules/concepts/hooks/list-of-hooks/displaypaymentreturn/) | paymentReturn | Payment return | display |
| FO |  | [displayPaymentTop](/9/modules/concepts/hooks/list-of-hooks/displaypaymenttop/) | paymentTop | Top of payment page | display |
| FO  BO |  | [displayPDF<Template>](/9/modules/concepts/hooks/list-of-hooks/displaypdftemplate/) |  | Add content to PDF templates | display |
| FO |  | [displayPersonalInformationTop](/9/modules/concepts/hooks/list-of-hooks/displaypersonalinformationtop/) |  | Content in the checkout funnel, on top of the personal information panel | display |
| FO |  | [displayProductActions](/9/modules/concepts/hooks/list-of-hooks/displayproductactions/) |  | Display additional action button on the product page | action |
| FO |  | [displayProductAdditionalInfo](/9/modules/concepts/hooks/list-of-hooks/displayproductadditionalinfo/) | productActions, displayProductButtons | Product page additional info | display |
| FO |  | [displayProductExtraContent](/9/modules/concepts/hooks/list-of-hooks/displayproductextracontent/) |  | Add content to the product page | display |
| FO |  | [displayProductListReviews](/9/modules/concepts/hooks/list-of-hooks/displayproductlistreviews/) |  | displayProductListReviews | display |
| FO |  | [displayProductPriceBlock](/9/modules/concepts/hooks/list-of-hooks/displayproductpriceblock/) |  | displayProductPriceBlock | display |
| FO |  | [displayReassurance](/9/modules/concepts/hooks/list-of-hooks/displayreassurance/) |  | displayReassurance | display |
| FO |  | [displayRightColumn](/9/modules/concepts/hooks/list-of-hooks/displayrightcolumn/) | extraRight | New elements on a page (right column) | display |
| FO |  | [displayRightColumnProduct](/9/modules/concepts/hooks/list-of-hooks/displayrightcolumnproduct/) | extraRight | New elements on the product page (right column) | display |
| FO |  | [displaySearch](/9/modules/concepts/hooks/list-of-hooks/displaysearch/) |  |  | display |
| FO |  | [displayShoppingCart](/9/modules/concepts/hooks/list-of-hooks/displayshoppingcart/) | shoppingCartExtra | Shopping cart - Additional button | display |
| FO |  | [displayShoppingCartFooter](/9/modules/concepts/hooks/list-of-hooks/displayshoppingcartfooter/) | shoppingCart | Shopping cart footer | display |
| FO |  | [displayTop](/9/modules/concepts/hooks/list-of-hooks/displaytop/) | top | Top of pages | display |
| FO |  | [displayWrapperBottom](/9/modules/concepts/hooks/list-of-hooks/displaywrapperbottom/) |  | Main wrapper section (bottom) | display |
| FO |  | [displayWrapperTop](/9/modules/concepts/hooks/list-of-hooks/displaywrappertop/) |  | Main wrapper section (top) | display |
| FO |  | [filterCategoryContent](/9/modules/concepts/hooks/list-of-hooks/filtercategorycontent/) |  | Filter the content page category |  |
| FO |  | [filterCmsCategoryContent](/9/modules/concepts/hooks/list-of-hooks/filtercmscategorycontent/) |  | Filter the content page category |  |
| FO |  | [filterCmsContent](/9/modules/concepts/hooks/list-of-hooks/filtercmscontent/) |  | Filter the content page |  |
| FO |  | [filterHtmlContent](/9/modules/concepts/hooks/list-of-hooks/filterhtmlcontent/) |  | Filter HTML field before rending a page |  |
| FO |  | [filterManufacturerContent](/9/modules/concepts/hooks/list-of-hooks/filtermanufacturercontent/) |  | Filter the content page manufacturer |  |
| FO |  | [filterProductContent](/9/modules/concepts/hooks/list-of-hooks/filterproductcontent/) |  | Filter the content page product |  |
| FO |  | [filterProductSearch](/9/modules/concepts/hooks/list-of-hooks/filterproductsearch/) |  | Filter search products result |  |
| FO |  | [filterSupplierContent](/9/modules/concepts/hooks/list-of-hooks/filtersuppliercontent/) |  | Filter the content page supplier |  |
| FO |  | [gSitemapAppendUrls](/9/modules/concepts/hooks/list-of-hooks/gsitemapappendurls/) |  |  |  |
| BO |  | [legacyblockkpi](/9/modules/concepts/hooks/list-of-hooks/legacyblockkpi/) |  |  |  |
| FO | ✅ | [moduleRoutes](/9/modules/concepts/hooks/list-of-hooks/moduleroutes/) |  |  |  |
| FO |  | [overrideImageLink](/9/modules/concepts/hooks/list-of-hooks/overrideimagelink/) |  | Override product image link | action |
| FO |  | [overrideLayoutTemplate](/9/modules/concepts/hooks/list-of-hooks/overridelayouttemplate/) |  |  |  |
| FO |  | [overrideMinimalPurchasePrice](/9/modules/concepts/hooks/list-of-hooks/overrideminimalpurchaseprice/) |  |  |  |
| FO |  | [productSearchProvider](/9/modules/concepts/hooks/list-of-hooks/productsearchprovider/) |  |  |  |
| FO |  | [sendMailAlterTemplateVars](/9/modules/concepts/hooks/list-of-hooks/sendmailaltertemplatevars/) |  | Alter template vars on the fly |  |
| FO |  | [termsAndConditions](/9/modules/concepts/hooks/list-of-hooks/termsandconditions/) |  |  |  |
| FO |  | [validateCustomerFormFields](/9/modules/concepts/hooks/list-of-hooks/validatecustomerformfields/) |  | Customer registration form validation |  |

## Related in KB

- [[hooks-overview|Hooks overview]]
- [[../themes/theme-hooks|Theme hooks]]
- [[../../modules/stmegamenu|stmegamenu uses displayTop]]
