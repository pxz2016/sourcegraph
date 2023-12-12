import { type FC, useEffect, useMemo } from 'react'

import type { TelemetryProps } from '@sourcegraph/shared/src/telemetry/telemetryService'
import {
    Link,
    PageHeader,
    useObservable,
    FORM_ERROR,
    type FormChangeEvent,
    type SubmissionErrors,
} from '@sourcegraph/wildcard'

import { PageTitle } from '../../../../../../components/PageTitle'
import { CodeInsightsIcon } from '../../../../../../insights/Icons'
import { CodeInsightCreationMode, CodeInsightsCreationActions, CodeInsightsPage } from '../../../../components'
import type { MinimalCaptureGroupInsightData } from '../../../../core'
import { useUiFeatures } from '../../../../hooks'
import { CodeInsightTrackType } from '../../../../pings'

import { CaptureGroupCreationContent } from './components/CaptureGroupCreationContent'
import { useCaptureInsightInitialValues } from './hooks/use-capture-insight-initial-values'
import type { CaptureGroupFormFields } from './types'
import { getSanitizedCaptureGroupInsight } from './utils/capture-group-insight-sanitizer'

interface CaptureGroupCreationPageProps extends TelemetryProps {
    backUrl: string
    onInsightCreateRequest: (event: { insight: MinimalCaptureGroupInsightData }) => Promise<unknown>
    onSuccessfulCreation: () => void
    onCancel: () => void
}

export const CaptureGroupCreationPage: FC<CaptureGroupCreationPageProps> = props => {
    const { backUrl, telemetryService, telemetryRecorder, onInsightCreateRequest, onSuccessfulCreation, onCancel } =
        props

    const { licensed, insight } = useUiFeatures()
    const creationPermission = useObservable(useMemo(() => insight.getCreationPermissions(), [insight]))

    const [initialFormValues, setInitialFormValues] = useCaptureInsightInitialValues()

    useEffect(() => {
        telemetryService.logViewEvent('CodeInsightsCaptureGroupCreationPage')
        telemetryRecorder.recordEvent('codeInsightsCaptureGroupCreationPage', 'viewed')
    }, [telemetryService, telemetryRecorder])

    const handleSubmit = async (values: CaptureGroupFormFields): Promise<SubmissionErrors | void> => {
        const insight = getSanitizedCaptureGroupInsight(values)

        await onInsightCreateRequest({ insight })

        setInitialFormValues(undefined)
        telemetryService.log('CodeInsightsCaptureGroupCreationPageSubmitClick')
        telemetryRecorder.recordEvent('codeInsightsCaptureGroupCreationPageSubmit', 'clicked')
        telemetryService.log(
            'InsightAddition',
            { insightType: CodeInsightTrackType.CaptureGroupInsight },
            { insightType: CodeInsightTrackType.CaptureGroupInsight }
        )
        telemetryRecorder.recordEvent('insightAddition', 'added', {
            privateMetadata: { insightType: CodeInsightTrackType.CaptureGroupInsight },
        })

        onSuccessfulCreation()
    }

    const handleCancel = (): void => {
        // Clear initial values if user successfully created search insight
        setInitialFormValues(undefined)
        telemetryService.log('CodeInsightsCaptureGroupCreationPageCancelClick')
        telemetryRecorder.recordEvent('codeInsightsCaptureGroupCreationPageCancel', 'clicked')

        onCancel()
    }

    const handleChange = (event: FormChangeEvent<CaptureGroupFormFields>): void => {
        setInitialFormValues(event.values)
    }

    return (
        <CodeInsightsPage>
            <PageTitle title="Create detect and track patterns insight - Code Insights" />

            <PageHeader
                className="mb-5"
                path={[
                    { icon: CodeInsightsIcon, to: '/insights', ariaLabel: 'Code insights dashboard page' },
                    { text: 'Create', to: backUrl },
                    { text: 'Detect and track patterns insight' },
                ]}
                description={
                    <span className="text-muted">
                        Capture group code insights analyze your code based on generated data series queries.{' '}
                        <Link
                            to="/help/code_insights/explanations/automatically_generated_data_series"
                            target="_blank"
                            rel="noopener"
                        >
                            Learn more.
                        </Link>
                    </span>
                }
            />

            <CaptureGroupCreationContent
                touched={false}
                initialValues={initialFormValues}
                className="pb-5"
                onSubmit={handleSubmit}
                onCancel={handleCancel}
                onChange={handleChange}
            >
                {form => (
                    <CodeInsightsCreationActions
                        mode={CodeInsightCreationMode.Creation}
                        licensed={licensed}
                        available={creationPermission?.available}
                        submitting={form.submitting}
                        errors={form.submitErrors?.[FORM_ERROR]}
                        clear={form.isFormClearActive}
                        onCancel={handleCancel}
                    />
                )}
            </CaptureGroupCreationContent>
        </CodeInsightsPage>
    )
}
